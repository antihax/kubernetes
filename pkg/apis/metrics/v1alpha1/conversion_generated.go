// +build !ignore_autogenerated

/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-conversions.sh

package v1alpha1

import (
	reflect "reflect"

	api "k8s.io/kubernetes/pkg/api"
	metrics "k8s.io/kubernetes/pkg/apis/metrics"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func autoConvert_metrics_RawNode_To_v1alpha1_RawNode(in *metrics.RawNode, out *RawNode, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*metrics.RawNode))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_metrics_RawNode_To_v1alpha1_RawNode(in *metrics.RawNode, out *RawNode, s conversion.Scope) error {
	return autoConvert_metrics_RawNode_To_v1alpha1_RawNode(in, out, s)
}

func autoConvert_metrics_RawPod_To_v1alpha1_RawPod(in *metrics.RawPod, out *RawPod, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*metrics.RawPod))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_metrics_RawPod_To_v1alpha1_RawPod(in *metrics.RawPod, out *RawPod, s conversion.Scope) error {
	return autoConvert_metrics_RawPod_To_v1alpha1_RawPod(in, out, s)
}

func autoConvert_v1alpha1_RawNode_To_metrics_RawNode(in *RawNode, out *metrics.RawNode, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*RawNode))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_RawNode_To_metrics_RawNode(in *RawNode, out *metrics.RawNode, s conversion.Scope) error {
	return autoConvert_v1alpha1_RawNode_To_metrics_RawNode(in, out, s)
}

func autoConvert_v1alpha1_RawPod_To_metrics_RawPod(in *RawPod, out *metrics.RawPod, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*RawPod))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_RawPod_To_metrics_RawPod(in *RawPod, out *metrics.RawPod, s conversion.Scope) error {
	return autoConvert_v1alpha1_RawPod_To_metrics_RawPod(in, out, s)
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		autoConvert_metrics_RawNode_To_v1alpha1_RawNode,
		autoConvert_metrics_RawPod_To_v1alpha1_RawPod,
		autoConvert_v1alpha1_RawNode_To_metrics_RawNode,
		autoConvert_v1alpha1_RawPod_To_metrics_RawPod,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}
