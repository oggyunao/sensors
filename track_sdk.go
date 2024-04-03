/*
 * Created by dengshiwei on 2022/06/06.
 * Copyright 2015－2022 Sensors Data Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sensorsanalytics

import (
	"encoding/json"
	"math/rand"
	"time"

	"oasis/pkg/sensors/v2.1.3/structs"
	"oasis/pkg/sensors/v2.1.3/utils"
)

func TrackEvent2(sa *SensorsAnalytics, etype, event, distinctId, originId string, properties map[string]interface{}, isLoginId bool) ([]byte, error) {
	eventTime := utils.NowMs()
	if properties == nil {
		properties = map[string]interface{}{}
	}
	if et := extractUserTime(properties); et > 0 {
		eventTime = et
	}
	rand.Seed(time.Now().UnixNano())
	data := structs.EventData{
		Type:          etype,
		TrackID:       rand.Int31(),
		Time:          eventTime,
		DistinctId:    distinctId,
		Properties:    properties,
		LibProperties: getLibProperties(),
	}

	project := getProject(data.Properties, sa.ProjectName)
	if project != "" {
		data.Project = project
	}

	if etype == TRACK || etype == TRACK_SIGNUP {
		data.Event = event
		properties["$lib"] = LIB_NAME
		properties["$lib_version"] = SDK_VERSION
	}

	if etype == TRACK_SIGNUP {
		data.OriginId = originId
	}

	if sa.TimeFree {
		data.TimeFree = true
	}

	if isLoginId {
		properties["$is_login_id"] = true
	}

	err := data.NormalizeData()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

func ItemTrack2(sa *SensorsAnalytics, trackType string, itemType string, itemId string, properties map[string]interface{}) ([]byte, error) {
	eventTime := utils.NowMs()
	if et := extractUserTime(properties); et > 0 {
		eventTime = et
	}
	libProperties := getLibProperties()
	var nproperties map[string]interface{}
	// merge properties
	if properties == nil {
		nproperties = make(map[string]interface{})
	} else {
		nproperties = utils.DeepCopy(properties)
	}

	rand.Seed(time.Now().UnixNano())
	itemData := structs.Item{
		Type:          trackType,
		ItemId:        itemId,
		TrackID:       rand.Int(),
		Time:          eventTime,
		ItemType:      itemType,
		Properties:    nproperties,
		LibProperties: libProperties,
	}
	project := getProject(itemData.Properties, sa.ProjectName)
	if project != "" {
		itemData.Project = project
	}
	err := itemData.NormalizeItem()
	if err != nil {
		return nil, err
	}

	return json.Marshal(itemData)
}

func TrackEventID32(sa *SensorsAnalytics, identity Identity, etype, event string, properties map[string]interface{}) ([]byte, error) {
	eventTime := utils.NowMs()
	if properties == nil {
		properties = map[string]interface{}{}
	}

	if et := extractUserTime(properties); et > 0 {
		eventTime = et
	}
	rand.Seed(time.Now().UnixNano())
	data := structs.EventData{
		Type:          etype,
		TrackID:       rand.Int31(),
		Time:          eventTime,
		Identities:    identity.Identities,
		Properties:    properties,
		LibProperties: getLibProperties(),
	}

	err := data.CheckIdentities()
	if err != nil {
		return nil, err
	}

	// 添加 distinct_id
	var distinctId string
	idValue := identity.Identities[LOGIN_ID]
	if len(idValue) <= 0 {
		for k, v := range identity.Identities {
			distinctId = k + "+" + v
		}
	} else {
		distinctId = idValue
	}
	data.DistinctId = distinctId

	project := getProject(data.Properties, sa.ProjectName)
	if project != "" {
		data.Project = project
	}

	if etype == TRACK || etype == BIND || etype == UNBIND {
		data.Event = event
		properties["$lib"] = LIB_NAME
		properties["$lib_version"] = SDK_VERSION
	}

	if sa.TimeFree {
		data.TimeFree = true
	}

	err = data.NormalizeData()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
