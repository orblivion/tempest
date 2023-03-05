package browsermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/util"
	"zenhack.net/go/util/maybe"
	"zenhack.net/go/util/orerr"
)

type Cmd = func(context.Context, chan<- Msg)

type Msg interface {
	Apply(Model) (Model, Cmd)
}

type UpsertGrain struct {
	Id    ID[Grain]
	Grain Grain
}

type RemoveGrain struct {
	Id ID[Grain]
}

type ClearGrains struct{}

type UpsertPackage struct {
	Id  ID[external.Package]
	Pkg external.Package
}

type RemovePackage struct {
	Id ID[external.Package]
}

type ClearPackages struct{}

type ChangeFocus struct {
	NewFocus Focus
}

type FocusGrain struct {
	Id ID[Grain]
}

type SpawnGrain struct {
	PkgID ID[external.Package]
}

type LoginSessionResult struct {
	Result orerr.T[external.LoginSession]
}

func (msg UpsertGrain) Apply(m Model) (Model, Cmd) {
	m.Grains[msg.Id].Handle.Release()
	m.Grains[msg.Id] = msg.Grain
	return m, nil
}

func (msg RemoveGrain) Apply(m Model) (Model, Cmd) {
	m.Grains[msg.Id].Handle.Release()
	delete(m.Grains, msg.Id)
	return m, nil
}

func (ClearGrains) Apply(m Model) (Model, Cmd) {
	m.Grains = make(map[ID[Grain]]Grain)
	return m, nil
}

func (msg UpsertPackage) Apply(m Model) (Model, Cmd) {
	m.Packages[msg.Id].Controller().Release()
	m.Packages[msg.Id] = msg.Pkg
	return m, nil
}

func (msg RemovePackage) Apply(m Model) (Model, Cmd) {
	// TODO(perf): release the whole message?
	m.Packages[msg.Id].Controller().Release()
	delete(m.Packages, msg.Id)
	return m, nil
}

func (ClearPackages) Apply(m Model) (Model, Cmd) {
	m.Packages = make(map[ID[external.Package]]external.Package)
	return m, nil
}

func (msg ChangeFocus) Apply(m Model) (Model, Cmd) {
	m.CurrentFocus = msg.NewFocus
	return m, nil
}

func (msg FocusGrain) Apply(m Model) (Model, Cmd) {
	m.CurrentFocus = FocusOpenGrain
	m.FocusedGrain = msg.Id
	_, ok := m.OpenGrains[msg.Id]
	if !ok {
		index := m.GrainDomOrder.Add(msg.Id)
		m.OpenGrains[msg.Id] = OpenGrain{
			DomainNonce: newDomainNonce(),
			DomIndex:    index,
		}
	}
	return m, nil
}

func (msg SpawnGrain) Apply(m Model) (Model, Cmd) {
	pkg := m.Packages[msg.PkgID]

	ctrl := pkg.Controller().AddRef()

	return m, func(ctx context.Context, msgs chan<- Msg) {
		defer ctrl.Release()
		fut, rel := ctrl.Create(ctx, func(p external.Package_Controller_create_Params) error {
			// TODO: pick something better, looking at the action/manifest:
			p.SetTitle("Untitled Grain")
			// TODO: provide a way to choose this:
			p.SetActionIndex(0)
			return nil
		})
		defer rel()
		res, err := fut.Struct()
		util.Chkfatal(err) // FIXME: report this properly somehow.

		id, err := res.Id()
		util.Chkfatal(err) // FIXME
		grain, err := res.Grain()
		util.Chkfatal(err) // FIXME

		title, err := grain.Title()
		util.Chkfatal(err) // FIXME
		sessionToken, err := grain.SessionToken()
		util.Chkfatal(err) // FIXME

		msgs <- UpsertGrain{
			Id: ID[Grain](id),
			Grain: Grain{
				Title:        title,
				SessionToken: sessionToken,
				Handle:       grain.Handle().AddRef(),
			},
		}
		msgs <- FocusGrain{Id: ID[Grain](id)}
	}
}

func (msg LoginSessionResult) Apply(m Model) (Model, Cmd) {
	m.LoginSession = maybe.New(msg.Result)
	sess, err := msg.Result.Get()
	if err != nil {
		return m, nil
	}
	return m, func(ctx context.Context, uiMsgs chan<- Msg) {
		pusher := collection.Pusher_ServerToClient(pkgPusher{
			uiMsgs: uiMsgs,
		})
		ret, rel := sess.ListPackages(context.Background(), func(p external.LoginSession_listPackages_Params) error {
			p.SetInto(pusher)
			return nil
		})
		defer rel()
		_, err := ret.Struct()
		if err != nil {
			println("listPackages(): " + err.Error())
		}
	}
}
