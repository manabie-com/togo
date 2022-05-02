package env_config

import (
	"os"
	"reflect"
	"strconv"
	"time"
)

func EnvStruct(s interface{}) error {
	out := Map(s)
	fields := structVal(s)
	for fieldName, env := range out {
		field := fields.FieldByName(fieldName)
		switch env.Value.(type) {
		case string:
			field.SetString(Env(env.Environment, env.DefaultValue))
		case int:
			data, err := EnvInt(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case int8:
			data, err := EnvInt8(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case int16:
			data, err := EnvInt16(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case int32:
			data, err := EnvInt32(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case int64:
			data, err := EnvInt64(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case uint:
			data, err := EnvUint(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case uint8:
			data, err := EnvUint8(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case uint16:
			data, err := EnvUint16(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case uint32:
			data, err := EnvUint32(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case uint64:
			data, err := EnvUint64(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case float32:
			data, err := EnvFloat32(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case float64:
			data, err := EnvFloat64(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case []byte:
			data := EnvBytes(env.Environment, env.DefaultValue)
			field.Set(reflect.ValueOf(data))
		case bool:
			data, err := EnvBool(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case time.Duration:
			data, err := EnvDuration(env.Environment, env.DefaultValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		case time.Time:
			data, err := EnvTime(env.Environment, env.DefaultValue, "")
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(data))
		}
	}
	return nil
}

// Env String
func Env(key, defaultValue string) (value string) {
	if value = os.Getenv(key); value == "" {
		value = defaultValue
	}
	return
}

// EnvBytes Byte
func EnvBytes(key, defaultValue string) []byte {
	return []byte(Env(key, defaultValue))
}

// EnvDuration time Duration
func EnvDuration(key, defaultValue string) (time.Duration, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	return time.ParseDuration(env)
}

// EnvTime time
func EnvTime(key, defaultValue, layout string) (time.Time, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return time.Time{}, nil
	}
	if layout == "" {
		layout = time.RFC3339
	}
	return time.Parse(layout, env)
}

// EnvInt Integer
func EnvInt(key, defaultValue string) (int, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	return strconv.Atoi(env)
}

func EnvInt8(key, defaultValue string) (int8, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseInt(env, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), nil
}

func EnvInt16(key, defaultValue string) (int16, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseInt(env, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), nil
}

func EnvInt32(key, defaultValue string) (int32, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseInt(env, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func EnvInt64(key, defaultValue string) (int64, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	return strconv.ParseInt(env, 10, 64)
}

// EnvUint Uint
func EnvUint(key, defaultValue string) (uint, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseUint(env, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func EnvUint8(key, defaultValue string) (uint8, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseUint(env, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(value), nil
}

func EnvUint16(key, defaultValue string) (uint16, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseUint(env, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(value), nil
}

func EnvUint32(key, defaultValue string) (uint32, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseUint(env, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(value), nil
}

func EnvUint64(key, defaultValue string) (uint64, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	return strconv.ParseUint(env, 10, 64)
}

// EnvFloat32 Float
func EnvFloat32(key, defaultValue string) (float32, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	value, err := strconv.ParseFloat(env, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

func EnvFloat64(key, defaultValue string) (float64, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return 0, nil
	}
	return strconv.ParseFloat(env, 64)
}

// EnvBool Boolean
func EnvBool(key, defaultValue string) (bool, error) {
	env := Env(key, defaultValue)
	if env == "" {
		return false, nil
	}
	return strconv.ParseBool(env)
}