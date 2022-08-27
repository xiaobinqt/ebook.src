package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-martini/martini"
)

type Logger interface {
	LogResources(resources R) Logger
	LogExtra(ex Extra) Logger
	LogError(err error) Logger
	AsyncFinish()
}

type (
	R     map[string][]interface{}
	Extra map[string]interface{}
)

const (
	CateAPIKey  = "api_keys"
	CateCerts   = "certs"
	CateRecip   = "recipients"
	CateSvrPool = "svrpools"

	CateImage        = "images"
	CateAlarm        = "alarm"
	CateRegistry     = "registry"
	CateAPP          = "apps"
	CateSetting      = "settings"
	CateNode         = "nodes"
	CateUser         = "users"
	CateContainer    = "containers"
	CateCCFG         = "ccfg"
	CateMigrate      = "migrate"
	CateBuildPack    = "buildpack"
	CateAppStore     = "app_store"
	CateAgentVersion = "agent_version"
	CateVolume       = "volume"
	CateLB           = "lb"
	CatePool         = "pool"
	CateCI           = "ci"
	CateTenant       = "tenant"
)

const (
	ActCreate       = "create"
	ActAdd          = "add"
	ActSub          = "sub"
	ActDelete       = "remove"
	ActUpdate       = "update"
	ActReset        = "reset"
	ActUpgrade      = "upgrade"
	ActTrigger      = "trigger"
	ActKill         = "kill"
	ActStart        = "start"
	ActStop         = "stop"
	ActRestart      = "restart"
	ActPause        = "pause"
	ActUnpause      = "unpause"
	ActCopy         = "copy"
	ActExec         = "exec"
	ActDeplay       = "deploy"
	ActRedeploy     = "redeploy"
	ActChangePasswd = "change_password"
	ActRestore      = "restore"
	ActMaintain     = "maintain"
	ActMove         = "move"
	ActBind         = "bind"
	ActUnbind       = "unbind"
	ActCommit       = "commit"
	ActView         = "view"
	ActRollback     = "rollback"

	ActAddNotification   = "add_notification"
	ActDistrubute        = "distrubute"
	ActAddTemplate       = "add_template"
	ActDelTemplate       = "delete_template"
	ActAddDescription    = "add_description"
	ActUpdateDescription = "update_description"
	ActAddInstance       = "add_instance"
	ActDelService        = "delete_service"
	ActCreateAPIKey      = "create_api_key"
	ActDeleteAPIKey      = "create_delete_key"
	ActGetLogs           = "get_logs"
	ActPush              = "push"
	ActAddTag            = "add_tags"
	ActDeleteTag         = "delete_tag"
	ActProxy             = "proxy"
	ActRename            = "rename"
	ActMigrate           = "migrate"

	ActSetSmtp       = "set_smtp"
	ActUpdateSmtp    = "update_smtp"
	ActDelRegImgTag  = "del_reg_img_tag"
	ActUpdateLicense = "update_license"
	ActDontTrack     = "dont_track"
	ActUpdateQuota   = "update_quota"

	ActPubVariables  = "pub_servicevars"
	ActApplyRevision = "apply_revisions"
	ActChangSum      = "chang_sum"
	ActAddContainer  = "add_container"
	ActDelContainer  = "del_container"

	// Legacy
	ActAddCcfgTemplate = "add_ccfgtemplate"
	ActDelCcfgTemplate = "del_ccfgtemplate"
	ActAddCcfgRevision = "add_ccfgrevision"
	ActPubCcfgRevision = "pub_ccfgrevision"

	ActMigrateExport = "migrate_export"
	ActMigrateImport = "migrate_import"

	ActBindBalance   = "bind_balance"
	ActUpdateBalance = "update_balance"
	ActUnbindBalance = "unbind_balance"
	ActSwitchBalance = "switch_balance"

	ActAddCIProject    = "add_ci_project"
	ActUpdateCIProject = "update_ci_project"
	ActDelCIProject    = "del_ci_project"
	ActStartCIBuild    = "start_ci_build"
	ActCancelCIBuild   = "cancel_ci_build"

	ActSwitchTenant = "switch_tenant"
)

const (
	RKeyJob           = "job"
	RKeyTemplate      = "template"
	RKeyAppTemplate   = "app_template"
	RKeyContainer     = "container"
	RKeyContainerName = "container_name"
	RKeyAlarm         = "alarm"
	RKeyAlarmPolicy   = "alarm_policy"
	RKeyRecipient     = "recipient"
	RKeySetting       = "setting"
	RKeyInstance      = "instance"
	RKeyService       = "instance_service"
	RKeyUser          = "user"
	RKeyLicense       = "license"
	RKeyAPIKey        = "api_key"
	RKeyNode          = "node"
	RKeyNodes         = "nodes"
	RKeySvrPool       = "svrpool"
	RKeyRegistry      = "registry"
	RKeyRegNoti       = "registry_notification"
	RKeyExec          = "exec"
	RKeyAutoBuild     = "auto_build"
	RKeyImage         = "image"
	RKeyImageID       = "image_id"
	RKeyResourceID    = "resource_id"
	RKeyTag           = "tag"
	RKeyQuota         = "quota"
	RKeyBuildPack     = "buildpack"
	RKeyRepository    = "repository" // git or svn repository
	RKeyRevision      = "revision"
	RKeyNewRevision   = "new_revision"
	RKeyDiffSum       = "diffsum"
	RKeyLb            = "lb"
	RKeyPool          = "pool"

	// Legacy
	RKeyCcfgTemplate     = "ccfg_template"
	RKeyCcfgRevision     = "ccfg_revision"
	RKeyCcfgTemplateName = "ccfg_templatename"

	// CI
	RkeyCIProject = "ci_project"
	RKeyCIBuild   = "ci_build"

	RKeyTenant = "tenant"
)

const (
	ExKeyError     = "error"
	ExKeyErrString = "error_string"
	ExKeyModify    = "modify"
	ExKeyPublish   = "publish"
)

const (
	OpStatusOK           = "successful"
	OpStatusFailed       = "failed"
	OpStatusUnknown      = "unknown"
	OpStatusSkipRedeploy = "skip"
)

type Entry struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Category   string        `json:"category" bson:"category"`
	Status     string        `json:"status" bson:"status"`
	Action     string        `json:"action" bson:"action"`
	StartAt    time.Time     `json:"started_at" bson:"started_at"`
	FinishedAt time.Time     `json:"finished_at" bson:"finished_at"`
	Operator   string        `json:"operator" bson:"operator"`
	RemoteAddr string        `json:"remote_addr" bson:"remote_addr"`
	Version    string        `json:"version" bson:"version"`
	ErrorCode  int           `json:"error_code" bson:"error_code"`
	Resources  R             `json:"resources" bson:"resources"`
	Extra      Extra         `json:"extra" bson:"extra"`

	logged bool
}

func (e *Entry) LogResources(resources R) Logger {
	if e.Resources == nil {
		e.Resources = resources
	} else {
		for k, v := range resources {
			e.Resources[k] = v
		}
	}

	e.logged = true
	return e
}

func (e *Entry) LogExtra(ex Extra) Logger {
	if e.Extra == nil {
		e.Extra = ex
	} else {
		for k, v := range ex {
			e.Extra[k] = v
		}
	}

	e.logged = true
	return e
}

func (e *Entry) LogError(err error) Logger {
	if err == nil {
		return e
	}
	return e.LogExtra(Extra{ExKeyError: err, ExKeyErrString: err.Error()})
}

func (e *Entry) AsyncFinish() {
	if !e.logged {
		return
	}
	e.FinishedAt = time.Now()
	if _, ok := e.Extra[ExKeyError]; ok {
		e.Status = OpStatusFailed
	} else {
		e.Status = OpStatusOK
	}

	fmt.Println("AsyncFinish...........")

	//if err := global.DB().Insert(global.C_AUDIT, e); err != nil {
	//	logrus.Errorf("failed to insert audit log %+v", e)
	//}
}

func NewEntry(category, action, operator, addr string) *Entry {
	return &Entry{
		ID:         bson.NewObjectId(),
		Category:   category,
		Action:     action,
		StartAt:    time.Now(),
		Operator:   operator,
		RemoteAddr: addr,
		Version:    "3.0.1",
	}
}

func LogHandler(category, action, operator string, w http.ResponseWriter, req *http.Request, c martini.Context,
	ctx *RequestContext) {
	addr := req.Header.Get("X-Real-IP")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
		}
	}
	e := NewEntry(category, action, operator, addr)

	c.MapTo(e, (*Logger)(nil))

	c.Next()

	if !e.logged || action == "" {
		return
	}
	e.FinishedAt = time.Now()
	e.ErrorCode = w.(martini.ResponseWriter).Status()
	if e.ErrorCode >= 400 {
		e.Status = OpStatusFailed
	} else {
		e.Status = OpStatusOK
	}

	fmt.Println("LogHandler...........", ctx.LogOldVal)
	//if err := global.DB().Insert(global.C_AUDIT, e); err != nil {
	//	logrus.Errorf("failed to insert audit log %+v", e)
	//}
}
