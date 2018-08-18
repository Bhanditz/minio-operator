/*
 * Minio-Operator - Manage Minio clusters in Kubernetes
 *
 * Minio (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MinioInstance is a specification for a Minio resource
type MinioInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinioInstanceSpec   `json:"spec"`
	Status MinioInstanceStatus `json:"status"`
}

// MinioInstanceSpec is the spec for a MinioInstance resource
type MinioInstanceSpec struct {
	// Version defines the MinioInstance Docker image version.
	Version string `json:"version"`
	// Replicas defines the number of Minio instances in a MinioInstance resource
	Replicas int32 `json:"replicas"`
	// If provided, use this secret as the credentials for MinioInstance resource
	// Otherwise Minio server creates dynamic credentials printed on Minio server startup banner
	// +optional
	CredsSecret *corev1.LocalObjectReference `json:"credsSecret,omitempty"`
	// VolumeClaimTemplate allows a user to specify how volumes inside a MinioInstance
	// +optional
	VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// If specified, affinity will define the pod's scheduling constraints
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// SSLSecret allows a user to specify custom CA certificate, and private key for group replication SSL.
	// +optional
	SSLSecret *corev1.LocalObjectReference `json:"sslSecret,omitempty"`
}

// MinioInstanceStatus is the status for a MinioInstance resource
type MinioInstanceStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MinioInstanceList is a list of MinioInstance resources
type MinioInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MinioInstance `json:"items"`
}
