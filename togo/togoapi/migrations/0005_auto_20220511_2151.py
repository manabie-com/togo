# Generated by Django 3.2.13 on 2022-05-11 13:51

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('togoapi', '0004_auto_20220510_0028'),
    ]

    operations = [
        migrations.AlterField(
            model_name='task',
            name='description',
            field=models.CharField(blank=True, max_length=100),
        ),
        migrations.AlterField(
            model_name='task',
            name='end_time',
            field=models.DateTimeField(blank=True),
        ),
        migrations.AlterField(
            model_name='task',
            name='start_time',
            field=models.DateTimeField(blank=True),
        ),
    ]
