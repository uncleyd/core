package rpc

import "strconv"

type Req struct {
	Data map[string]interface{}
}

type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

type RespJsonList struct {
	Code  int
	Count int
	Data  interface{}
}

func (r *Req) GetString(key string) string {
	if _, ok := r.Data[key]; !ok {
		return ""
	}
	return r.Data[key].(string)
}

func (r *Req) GetInt(key string, def ...int) (int, error) {
	if _, ok := r.Data[key]; !ok {
		return def[0], nil
	}

	strv := r.Data[key].(string)

	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}

	return strconv.Atoi(strv)
}

func (r *Req) GetInt64(key string, def ...int64) (int64, error) {
	if _, ok := r.Data[key]; !ok {
		return def[0], nil
	}
	strv := r.Data[key].(string)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(strv, 10, 64)
}

func (r *Req) GetFloat(key string, def ...float64) (float64, error) {
	if _, ok := r.Data[key]; !ok {
		return def[0], nil
	}
	strv := r.Data[key].(string)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(strv, 64)
}
