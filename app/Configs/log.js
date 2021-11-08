import path from 'path'

const config = {
    pathLogError: path.resolve(process.cwd(), "storages/logs/error_log.log"),
    pathLogWarning: path.resolve(process.cwd(), "storages/logs/warning_log.log"),
    pathLogAccess: path.resolve(process.cwd(), "storages/logs/access_log.log")
}

export default config
