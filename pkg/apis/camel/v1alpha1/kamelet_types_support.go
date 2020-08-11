/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetConditions --
func (in *KameletStatus) GetConditions() []v1.ResourceCondition {
	res := make([]v1.ResourceCondition, 0, len(in.Conditions))
	for _, c := range in.Conditions {
		res = append(res, c)
	}
	return res
}

// GetType --
func (c KameletCondition) GetType() string {
	return string(c.Type)
}

// GetStatus --
func (c KameletCondition) GetStatus() corev1.ConditionStatus {
	return c.Status
}

// GetLastUpdateTime --
func (c KameletCondition) GetLastUpdateTime() metav1.Time {
	return c.LastUpdateTime
}

// GetLastTransitionTime --
func (c KameletCondition) GetLastTransitionTime() metav1.Time {
	return c.LastTransitionTime
}

// GetReason --
func (c KameletCondition) GetReason() string {
	return c.Reason
}

// GetMessage --
func (c KameletCondition) GetMessage() string {
	return c.Message
}