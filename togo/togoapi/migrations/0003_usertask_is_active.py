# Generated by Django 3.2.13 on 2022-05-08 14:03

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('togoapi', '0002_auto_20220507_2313'),
    ]

    operations = [
        migrations.AddField(
            model_name='usertask',
            name='is_active',
            field=models.BooleanField(default=True),
        ),
    ]
