# Generated by Django 3.2.5 on 2022-06-22 11:44

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('Models', '0002_detail_user'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='detail',
            name='user',
        ),
    ]
