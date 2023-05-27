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
# Generated by Django 3.2.12 on 2023-05-27 10:03

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0009_config_mount_app_log_to_host'),
        ('region', '0002_auto_20201216_1213'),
    ]

    operations = [
        migrations.CreateModel(
            name='EgressSpec',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now=True)),
                ('replicas', models.IntegerField(default=1)),
                ('cpu_limit', models.CharField(max_length=16)),
                ('memory_limit', models.CharField(max_length=16)),
                ('wl_app', models.OneToOneField(on_delete=django.db.models.deletion.CASCADE, to='api.app')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='EgressRule',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now=True)),
                ('dst_port', models.IntegerField(verbose_name='目标端口')),
                ('host', models.CharField(max_length=128, verbose_name='目标主机')),
                (
                    'protocol',
                    models.CharField(choices=[('TCP', 'TCP'), ('UDP', 'UDP')], max_length=32, verbose_name='协议'),
                ),
                ('src_port', models.IntegerField(verbose_name='源端口')),
                ('service', models.CharField(max_length=128, verbose_name='服务名')),
                (
                    'spec',
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE, related_name='rules', to='region.egressspec'
                    ),
                ),
            ],
            options={
                'abstract': False,
            },
        ),
    ]
