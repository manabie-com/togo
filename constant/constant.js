`use strict`;

function define(obj, name, value) {
    Object.defineProperty(obj, name, {
        value:        value,
        enumerable:   true,
        writable:     false,
        configurable: false
    });
}

exports.taskLimits = {
    LIMIT_TASK_PER_DAY: 5
};

exports.responseFlags = {};

define(exports.responseFlags, "SUCCESS", 0);
define(exports.responseFlags, "FULL_TASKS", 1);
define(exports.responseFlags, "INVALID", 2);