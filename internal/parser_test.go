package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextToken(t *testing.T) {
	data = []byte(`/* renamed from: com.tencent.pb.profilecard.VaProfileGate */
/* compiled from: P */
public final class VaProfileGate {`)
	assert.Equal(t, nextToken(), "public")
	assert.Equal(t, nextToken(), "final")
	assert.Equal(t, nextToken(), "class")
}

func TestParseFieldMap(t *testing.T) {
	data = []byte(`= MessageMicro.initFieldMap(new int[]{8, 16, 26, 34, 40, 50, 58}, new String[]{"cmd", "code", "version", "url", "interv", "buff", "appname"}, new Object[]{3, 0, "", "", 0, "", ""}, SUpdateRsp.class);
        public final PBStringField appname = PBField.initString("");
        public final PBStringField buff = PBField.initString("");`)
	fields := parseFieldMap()
	for _, p := range fields {
		fmt.Println("name: ", p.Name, " number: ", p.Tag)
	}
	assert.Equal(t, nextToken(), "public")
}
