// Package sensorsanalytics /*
package sensorsanalytics

import (
	"errors"
	"github.com/oggyunao/sensorsanalytics/utils"
)

func (sa *SensorsAnalytics) BindGenSendBytes3(identity Identity) ([]byte, error) {
	if identity.Identities == nil || len(identity.Identities) < 2 {
		return nil, errors.New("identity is invalid")
	}
	return TrackEventID32(sa, identity, BIND, BIND_EVENT, nil)
}

func (sa *SensorsAnalytics) UnBindGenSendBytes3(identity Identity) ([]byte, error) {
	if identity.Identities == nil {
		return nil, errors.New("identity is nil")
	}
	return TrackEventID32(sa, identity, UNBIND, UNBIND_EVENT, nil)
}

func (sa *SensorsAnalytics) TrackByIdGenSendBytes3(identity Identity, event string, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	// merge properties
	if properties == nil {
		nproperties = make(map[string]interface{})
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	// merge super properties
	if superProperties != nil {
		utils.MergeSuperProperty(superProperties, nproperties)
	}

	return TrackEventID32(sa, identity, TRACK, event, nproperties)
}

func (sa *SensorsAnalytics) ProfileSetByIdGenSendBytes3(identity Identity, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	if properties == nil {
		return nil, errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID32(sa, identity, PROFILE_SET, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileSetOnceByIdGenSendBytes3(identity Identity, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	if properties == nil {
		return nil, errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID32(sa, identity, PROFILE_SET_ONCE, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileIncrementByIdGenSendBytes3(identity Identity, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	if properties == nil {
		return nil, errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID32(sa, identity, PROFILE_INCREMENT, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileAppendByIdGenSendBytes3(identity Identity, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	if properties == nil {
		return nil, errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID32(sa, identity, PROFILE_APPEND, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileUnsetByIdGenSendBytes3(identity Identity, properties map[string]interface{}) ([]byte, error) {
	var nproperties map[string]interface{}

	if properties == nil {
		return nil, errors.New("property should not be nil")
	} else {
		nproperties = utils.DeepCopy(properties)
	}
	return TrackEventID32(sa, identity, PROFILE_UNSET, "", nproperties)
}

func (sa *SensorsAnalytics) ProfileDeleteByIdGenSendBytes3(identity Identity) ([]byte, error) {
	nproperties := make(map[string]interface{})
	return TrackEventID32(sa, identity, PROFILE_DELETE, "", nproperties)
}
