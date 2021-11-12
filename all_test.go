package format

import "testing"

func Test(t *testing.T) {
	t.Run("TestFormat", TestFormat)
	t.Run("TestSetDict", TestSetDict)
	t.Run("TestAppendDict", TestAppendDict)
}
