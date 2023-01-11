# -*- coding: utf-8 -*-
"""
TencentBlueKing is pleased to support the open source community by making
蓝鲸智云 - PaaS 平台 (BlueKing - PaaS System) available.
Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.

We undertake not to change the open source license (MIT license) applicable
to the current version of the project delivered to anyone in the future.
"""
from blue_krill.data_types.enum import EnumField, StructuredEnum


class MetricsSeriesType(str, StructuredEnum):
    CURRENT = EnumField('current', '使用量')
    REQUEST = EnumField('request', '保留量')
    LIMIT = EnumField('limit', '配额上限')


class MetricsResourceType(str, StructuredEnum):
    MEM = EnumField('mem')
    CPU = EnumField('cpu')

    @classmethod
    def choices(cls):
        """Return original choice tuples with an extra value: __all__"""
        choices = super().get_django_choices()
        return choices + [('__all__', '__ALL__')]


BKMONITOR_PROMQL_TMPL = {
    'mem': {
        # 内存实际使用值
        'current': 'sum by(container_name)(container_memory_working_set_bytes{'
        'pod_name="%s",container_name!="POD",bcs_cluster_id="%s",bk_biz_id="%s"})',
        # 内存预留值
        'request': 'kube_pod_container_resource_requests_memory_bytes{pod="%s",bcs_cluster_id="%s",bk_biz_id="%s"}',
        # 内存上限值
        'limit': 'kube_pod_container_resource_limits_memory_bytes{pod="%s",bcs_cluster_id="%s",bk_biz_id="%s"}',
    },
    'cpu': {
        # CPU 实际使用值
        'current': 'sum by(container_name)(rate(container_cpu_usage_seconds_total{'
        'image!="",pod_name="%s",container_name!="POD",bcs_cluster_id="%s",bk_biz_id="%s"}[1m]))',
        # CPU 预留值
        'request': 'kube_pod_container_resource_requests_cpu_cores{pod="%s",bcs_cluster_id="%s",bk_biz_id="%s"}',
        # CPU 上限值
        'limit': 'kube_pod_container_resource_limits_cpu_cores{pod="%s",bcs_cluster_id="%s",bk_biz_id="%s"}',
    },
}
