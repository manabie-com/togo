from src.util import load_config
from src import create_app

if __name__ == '__main__':
    app_name = 'manabie-togo'
    config: dict = load_config()
    app = create_app(app_name)
    app_config = config.get('APP', {})
    app.run(
        # host=app_config.get('host', "localhost"),
        port=app_config.get('port', 5500),
        debug=app_config.get('debug', True)
    )


