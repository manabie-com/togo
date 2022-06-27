import importlib
import six

string_types = six.string_types


def config_str_to_obj(path, cfg):
    if isinstance(cfg, string_types):
        module = importlib.import_module(path, [cfg])
        return getattr(module, cfg)
    return cfg
