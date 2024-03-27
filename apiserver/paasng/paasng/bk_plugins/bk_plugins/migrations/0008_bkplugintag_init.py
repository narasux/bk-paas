# Generated by Django 3.2.12 on 2023-02-27 11:24

from django.db import migrations


def init_bkplugins_tag(apps, schema_editor):
    """初始化插件分类，插件必须设置分类后才能部署生产环境"""
    BkPluginTag = apps.get_model('bk_plugins', 'BkPluginTag')
    BkPluginTag.objects.get_or_create(code_name="OTHER", defaults={
        "name": "未分类",
        "priority": 1,  
    })

class Migration(migrations.Migration):

    dependencies = [
        ('bk_plugins', '0007_auto_20230227_1924'),
    ]

    operations = [migrations.RunPython(init_bkplugins_tag)]
