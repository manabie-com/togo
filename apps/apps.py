from django.apps import AppConfig
from django.utils.translation import gettext_lazy as _


class AppsConfig(AppConfig):
    label = 'Apps'
    name = 'apps'
    verbose_name = _('Apps')
    default_auto_field = 'django.db.models.AutoField'
