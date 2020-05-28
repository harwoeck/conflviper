package conflviper

import (
	"time"

	"github.com/harwoeck/confl"
	"github.com/spf13/viper"
)

type ConflViper struct {
	viper *viper.Viper
}

var _ confl.Confl = (*ConflViper)(nil)

func NewConflViper(viper *viper.Viper) *ConflViper {
	return &ConflViper{
		viper: viper,
	}
}

func (c *ConflViper) Source() interface{} {
	return c.viper
}

func (c *ConflViper) InConfig(key string) bool {
	return c.viper.InConfig(key)
}

func (c *ConflViper) IsSet(key string) bool {
	return c.viper.IsSet(key)
}

func (c *ConflViper) Sub(key string) confl.Confl {
	return &ConflViper{
		viper: c.viper.Sub(key),
	}
}

func (c *ConflViper) Get(key string) interface{} {
	return c.viper.Get(key)
}
func (c *ConflViper) GetDefault(key string, val interface{}) interface{} {
	if c.IsSet(key) {
		return c.Get(key)
	}
	return val
}

func (c *ConflViper) GetBool(key string) bool {
	return c.viper.GetBool(key)
}
func (c *ConflViper) GetBoolDefault(key string, val bool) bool {
	if c.IsSet(key) {
		return c.GetBool(key)
	}
	return val
}

func (c *ConflViper) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}
func (c *ConflViper) GetDurationDefault(key string, val time.Duration) time.Duration {
	if c.IsSet(key) {
		return c.GetDuration(key)
	}
	return val
}

func (c *ConflViper) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}
func (c *ConflViper) GetFloat64Default(key string, val float64) float64 {
	if c.IsSet(key) {
		return c.GetFloat64(key)
	}
	return val
}

func (c *ConflViper) GetInt(key string) int {
	return c.viper.GetInt(key)
}
func (c *ConflViper) GetIntDefault(key string, val int) int {
	if c.IsSet(key) {
		return c.GetInt(key)
	}
	return val
}

func (c *ConflViper) GetInt32(key string) int32 {
	return c.viper.GetInt32(key)
}
func (c *ConflViper) GetInt32Default(key string, val int32) int32 {
	if c.IsSet(key) {
		return c.GetInt32(key)
	}
	return val
}

func (c *ConflViper) GetInt64(key string) int64 {
	return c.viper.GetInt64(key)
}
func (c *ConflViper) GetInt64Default(key string, val int64) int64 {
	if c.IsSet(key) {
		return c.GetInt64(key)
	}
	return val
}

func (c *ConflViper) GetIntSlice(key string) []int {
	return c.viper.GetIntSlice(key)
}
func (c *ConflViper) GetIntSliceDefault(key string, val []int) []int {
	if c.IsSet(key) {
		return c.GetIntSlice(key)
	}
	return val
}

func (c *ConflViper) GetSizeInBytes(key string) uint {
	return c.viper.GetSizeInBytes(key)
}
func (c *ConflViper) GetSizeInBytesDefault(key string, val uint) uint {
	if c.IsSet(key) {
		return c.GetSizeInBytes(key)
	}
	return val
}

func (c *ConflViper) GetString(key string) string {
	return c.viper.GetString(key)
}
func (c *ConflViper) GetStringDefault(key string, val string) string {
	if c.IsSet(key) {
		return c.GetString(key)
	}
	return val
}

func (c *ConflViper) GetStringMap(key string) map[string]interface{} {
	return c.viper.GetStringMap(key)
}
func (c *ConflViper) GetStringMapDefault(key string, val map[string]interface{}) map[string]interface{} {
	if c.IsSet(key) {
		return c.GetStringMap(key)
	}
	return val
}

func (c *ConflViper) GetStringMapString(key string) map[string]string {
	return c.viper.GetStringMapString(key)
}
func (c *ConflViper) GetStringMapStringDefault(key string, val map[string]string) map[string]string {
	if c.IsSet(key) {
		return c.GetStringMapString(key)
	}
	return val
}

func (c *ConflViper) GetStringMapStringSlice(key string) map[string][]string {
	return c.viper.GetStringMapStringSlice(key)
}
func (c *ConflViper) GetStringMapStringSliceDefault(key string, val map[string][]string) map[string][]string {
	if c.IsSet(key) {
		return c.GetStringMapStringSlice(key)
	}
	return val
}

func (c *ConflViper) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}
func (c *ConflViper) GetStringSliceDefault(key string, val []string) []string {
	if c.IsSet(key) {
		return c.GetStringSlice(key)
	}
	return val
}

func (c *ConflViper) GetTime(key string) time.Time {
	return c.viper.GetTime(key)
}
func (c *ConflViper) GetTimeDefault(key string, val time.Time) time.Time {
	if c.IsSet(key) {
		return c.GetTime(key)
	}
	return val
}

func (c *ConflViper) GetUint(key string) uint {
	return c.viper.GetUint(key)
}
func (c *ConflViper) GetUintDefault(key string, val uint) uint {
	if c.IsSet(key) {
		return c.GetUint(key)
	}
	return val
}

func (c *ConflViper) GetUint32(key string) uint32 {
	return c.viper.GetUint32(key)
}
func (c *ConflViper) GetUint32Default(key string, val uint32) uint32 {
	if c.IsSet(key) {
		return c.GetUint32(key)
	}
	return val
}

func (c *ConflViper) GetUint64(key string) uint64 {
	return c.viper.GetUint64(key)
}
func (c *ConflViper) GetUint64Default(key string, val uint64) uint64 {
	if c.IsSet(key) {
		return c.GetUint64(key)
	}
	return val
}
