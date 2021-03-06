package konfig

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBind(t *testing.T) {
	t.Run(
		"panic invalid type",
		func(t *testing.T) {
			var s = Instance()
			require.Panics(t, func() { s.Bind(1) })
		},
	)

	t.Run(
		"valid type map",
		func(t *testing.T) {
			var s = Instance()
			var m = make(map[string]interface{})
			require.NotPanics(t, func() { s.Bind(m) })
		},
	)

	t.Run(
		"valid type struct",
		func(t *testing.T) {
			type testConfig struct {
				v string `konfig:"v"`
			}
			var s = Instance()
			var tc testConfig
			require.NotPanics(t, func() { s.Bind(tc) })
		},
	)

}

func TestSetStruct(t *testing.T) {

	t.Run(
		"valid type struct",
		func(t *testing.T) {
			type TestConfigSub struct {
				VV  string  `konfig:"vv"`
				TT  int     `konfig:"tt"`
				B   bool    `konfig:"bool"`
				F   float64 `konfig:"float64"`
				U   uint64  `konfig:"uint64"`
				I64 int64   `konfig:"int64"`
			}
			type TestConfig struct {
				V    string        `konfig:"v"`
				T    TestConfigSub `konfig:"sub"`
				SubT *TestConfigSub
			}

			var expectedConfig = TestConfig{
				V: "test",
				T: TestConfigSub{
					VV:  "test2",
					TT:  1,
					B:   true,
					F:   1.9,
					I64: 1,
				},
				SubT: &TestConfigSub{
					VV: "",
					TT: 2,
				},
			}

			Init(DefaultConfig())

			var tc TestConfig
			require.NotPanics(t, func() { Bind(tc) })

			var v = Values{
				"v":           "test",
				"sub.vv":      "test2",
				"sub.tt":      1,
				"subt.tt":     2,
				"sub.bool":    true,
				"sub.float64": 1.9,
				"sub.int64":   int64(1),
			}

			v.load(Values{
				"v": "a",
			}, instance())

			var configValue = Value().(TestConfig)
			require.Equal(t, "test", configValue.V)
			require.Equal(t, "test2", configValue.T.VV)
			require.Equal(t, 1, configValue.T.TT)
			require.Equal(t, 2, configValue.SubT.TT)
			require.Equal(t, true, configValue.T.B)
			require.Equal(t, 1.9, configValue.T.F)
			require.Equal(t, int64(1), configValue.T.I64)

			require.Equal(t, expectedConfig, Value())

			var vv = Values{
				"v":      "test",
				"sub.vv": "test2",
			}

			vv.load(v, instance())

			configValue = Value().(TestConfig)
			require.Equal(t, "test", configValue.V)
			require.Equal(t, "test2", configValue.T.VV)
			require.Equal(t, 0, configValue.T.TT)
			require.Equal(t, 0, configValue.SubT.TT)
		},
	)

	t.Run(
		"valid type struct",
		func(t *testing.T) {
			type TestConfigSub struct {
				VV string `konfig:"vv"`
				TT int    `konfig:"tt"`
			}
			type TestConfig struct {
				V    string        `konfig:"v"`
				T    TestConfigSub `konfig:"sub"`
				SubT *TestConfigSub
			}

			var expectedConfig = TestConfig{
				V: "test",
				T: TestConfigSub{
					VV: "bar",
					TT: 1,
				},
				SubT: &TestConfigSub{
					VV: "",
					TT: 2,
				},
			}

			Init(DefaultConfig())

			var tc TestConfig
			require.NotPanics(t, func() { Bind(tc) })

			Set("v", "test")
			Set("sub.vv", "bar")
			Set("sub.tt", 1)
			Set("subt.tt", 2)

			var configValue = Value().(TestConfig)
			require.Equal(t, "test", configValue.V)
			require.Equal(t, "bar", configValue.T.VV)
			require.Equal(t, 1, configValue.T.TT)
			require.Equal(t, 2, configValue.SubT.TT)

			require.Equal(t, expectedConfig, Value())

		},
	)

	t.Run(
		"valid type map",
		func(t *testing.T) {
			Init(DefaultConfig())

			var tc = make(map[string]interface{})
			require.NotPanics(t, func() { Bind(tc) })

			Set("v", "test")
			Set("sub.vv", "bar")
			Set("sub.tt", 1)
			Set("subt.tt", 2)

			var configValue = Value().(map[string]interface{})
			require.Equal(t, "test", configValue["v"])
			require.Equal(t, "bar", configValue["sub.vv"])
			require.Equal(t, 1, configValue["sub.tt"])
			require.Equal(t, 2, configValue["subt.tt"])
		},
	)

	t.Run(
		"valid type map",
		func(t *testing.T) {
			Init(DefaultConfig())

			var tc = make(map[string]interface{})
			require.NotPanics(t, func() { Bind(tc) })

			var v = Values{
				"v":       "test",
				"sub.vv":  "test2",
				"sub.tt":  1,
				"subt.tt": 2,
			}

			v.load(Values{}, instance())

			var configValue = Value().(map[string]interface{})
			require.Equal(t, "test", configValue["v"])
			require.Equal(t, "test2", configValue["sub.vv"])
			require.Equal(t, 1, configValue["sub.tt"])
			require.Equal(t, 2, configValue["subt.tt"])
		},
	)
}
