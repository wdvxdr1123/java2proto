package tencent.im.oidb.cmd0xf9b;

import com.tencent.mobileqq.pb.ByteStringMicro;
import com.tencent.mobileqq.pb.MessageMicro;
import com.tencent.mobileqq.pb.PBBytesField;
import com.tencent.mobileqq.pb.PBDoubleField;
import com.tencent.mobileqq.pb.PBField;
import com.tencent.mobileqq.pb.PBRepeatMessageField;
import com.tencent.mobileqq.pb.PBUInt32Field;

/* compiled from: P */
public final class oidb_cmd0xf9b {

    /* compiled from: P */
    public final class RspBody extends MessageMicro<RspBody> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{10}, new String[]{"locations"}, new Object[]{null}, RspBody.class);
        public final PBRepeatMessageField<Location> locations = PBField.initRepeatMessage(Location.class);
    }

    private oidb_cmd0xf9b() {
    }

    /* compiled from: P */
    public final class ReqBody extends MessageMicro<ReqBody> {
        static final MessageMicro.FieldMap __fieldMap__;
        public final PBUInt32Field coordinate = PBField.initUInt32(0);
        public final PBDoubleField lat = PBField.initDouble(0.0d);
        public final PBDoubleField lon = PBField.initDouble(0.0d);
        public final PBUInt32Field num = PBField.initUInt32(0);
        public final PBUInt32Field start = PBField.initUInt32(0);

        static {
            Double valueOf = Double.valueOf(0.0d);
            __fieldMap__ = MessageMicro.initFieldMap(new int[]{9, 17, 24, 32, 40}, new String[]{"lat", "lon", "num", "start", "coordinate"}, new Object[]{valueOf, valueOf, 0, 0, 0}, ReqBody.class);
        }
    }

    /* compiled from: P */
    public final class Location extends MessageMicro<Location> {
        static final MessageMicro.FieldMap __fieldMap__;
        public final PBBytesField address = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField area_id = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField city = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBUInt32Field distance = PBField.initUInt32(0);
        public final PBDoubleField lat = PBField.initDouble(0.0d);
        public final PBDoubleField lon = PBField.initDouble(0.0d);
        public final PBBytesField title = PBField.initBytes(ByteStringMicro.EMPTY);

        static {
            Double valueOf = Double.valueOf(0.0d);
            __fieldMap__ = MessageMicro.initFieldMap(new int[]{9, 17, 26, 34, 40, 50, 58}, new String[]{"lat", "lon", "title", "address", "distance", "city", "area_id"}, new Object[]{valueOf, valueOf, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, 0, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY}, Location.class);
        }
    }
}