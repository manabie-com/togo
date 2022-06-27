from django.apps import AppConfig
from django.utils.translation import gettext_lazy as _


class TestsConfig(AppConfig):
    label = 'Tests'
    name = 'tests'
    verbose_name = _('Tests')
