package tencent.im.oidb.cmd0x88d;

import com.tencent.mobileqq.pb.ByteStringMicro;
import com.tencent.mobileqq.pb.MessageMicro;
import com.tencent.mobileqq.pb.PBBytesField;
import com.tencent.mobileqq.pb.PBField;
import com.tencent.mobileqq.pb.PBInt64Field;
import com.tencent.mobileqq.pb.PBRepeatField;
import com.tencent.mobileqq.pb.PBRepeatMessageField;
import com.tencent.mobileqq.pb.PBUInt32Field;
import com.tencent.mobileqq.pb.PBUInt64Field;
import tencent.im.oidb.cmd0xef0.oidb_0xef0;

/* compiled from: P */
public final class oidb_0x88d {

    /* compiled from: P */
    public final class GroupCardPrefix extends MessageMicro<GroupCardPrefix> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{10, 18}, new String[]{"bytes_introduction", "rpt_bytes_prefix"}, new Object[]{ByteStringMicro.EMPTY, ByteStringMicro.EMPTY}, GroupCardPrefix.class);
        public final PBBytesField bytes_introduction = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBRepeatField<ByteStringMicro> rpt_bytes_prefix = PBField.initRepeat(PBBytesField.__repeatHelper__);
    }

    /* compiled from: P */
    public final class GroupExInfoOnly extends MessageMicro<GroupExInfoOnly> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16}, new String[]{"uint32_tribe_id", "uint32_money_for_add_group"}, new Object[]{0, 0}, GroupExInfoOnly.class);
        public final PBUInt32Field uint32_money_for_add_group = PBField.initUInt32(0);
        public final PBUInt32Field uint32_tribe_id = PBField.initUInt32(0);
    }

    /* compiled from: P */
    public final class GroupGeoInfo extends MessageMicro<GroupGeoInfo> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16, 24, 32, 40, 50, 56}, new String[]{"uint64_owneruin", "uint32_settime", "uint32_cityid", "int64_longitude", "int64_latitude", "bytes_geocontent", "uint64_poi_id"}, new Object[]{0L, 0, 0, 0L, 0L, ByteStringMicro.EMPTY, 0L}, GroupGeoInfo.class);
        public final PBBytesField bytes_geocontent = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBInt64Field int64_latitude = PBField.initInt64(0);
        public final PBInt64Field int64_longitude = PBField.initInt64(0);
        public final PBUInt32Field uint32_cityid = PBField.initUInt32(0);
        public final PBUInt32Field uint32_settime = PBField.initUInt32(0);
        public final PBUInt64Field uint64_owneruin = PBField.initUInt64(0);
        public final PBUInt64Field uint64_poi_id = PBField.initUInt64(0);
    }

    /* compiled from: P */
    public final class GroupHeadPortrait extends MessageMicro<GroupHeadPortrait> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 18, 24, 32, 42}, new String[]{"uint32_pic_cnt", "rpt_msg_info", "uint32_default_id", "uint32_verifying_pic_cnt", "rpt_msg_verifyingpic_info"}, new Object[]{0, null, 0, 0, null}, GroupHeadPortrait.class);
        public final PBRepeatMessageField<GroupHeadPortraitInfo> rpt_msg_info = PBField.initRepeatMessage(GroupHeadPortraitInfo.class);
        public final PBRepeatMessageField<GroupHeadPortraitInfo> rpt_msg_verifyingpic_info = PBField.initRepeatMessage(GroupHeadPortraitInfo.class);
        public final PBUInt32Field uint32_default_id = PBField.initUInt32(0);
        public final PBUInt32Field uint32_pic_cnt = PBField.initUInt32(0);
        public final PBUInt32Field uint32_verifying_pic_cnt = PBField.initUInt32(0);
    }

    /* compiled from: P */
    public final class GroupHeadPortraitInfo extends MessageMicro<GroupHeadPortraitInfo> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16, 24, 32, 40}, new String[]{"rpt_uint32_pic_id", "uint32_left_x", "uint32_left_y", "uint32_right_x", "uint32_right_y"}, new Object[]{0, 0, 0, 0, 0}, GroupHeadPortraitInfo.class);
        public final PBUInt32Field rpt_uint32_pic_id = PBField.initUInt32(0);
        public final PBUInt32Field uint32_left_x = PBField.initUInt32(0);
        public final PBUInt32Field uint32_left_y = PBField.initUInt32(0);
        public final PBUInt32Field uint32_right_x = PBField.initUInt32(0);
        public final PBUInt32Field uint32_right_y = PBField.initUInt32(0);
    }

    /* compiled from: P */
    public final class GroupInfo extends MessageMicro<GroupInfo> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16, 24, 32, 40, 48, 56, 64, 72, 80, 88, 96, 104, 112, 122, 130, 138, 146, 152, 160, 168, 176, 184, 194, 202, 208, 216, 224, 232, 240, 248, 258, 266, 272, 282, 288, 296, 304, 314, 322, 330, 338, 344, 354, 360, 368, 376, 384, 392, 400, 408, 416, 424, 432, 440, 448, 458, 464, 472, 480, 490, 496, 504, 512, 520, 528, 536, 544, 552, 560, 568, 576, 584, 592, 600, 608, 616, 624, 632, 640, 648, 656, 664, 672, 680, 688, 696, 704, 714, 720, 728, 736, 746, 752, 760, 768, 776, 784, 792, 800, 808, 816, 826, 832, 842, 848, 856, 864, 874, 882}, new String[]{"uint64_group_owner", "uint32_group_create_time", "uint32_group_flag", "uint32_group_flag_ext", "uint32_group_member_max_num", "uint32_group_member_num", "uint32_group_option", "uint32_group_class_ext", "uint32_group_special_class", "uint32_group_level", "uint32_group_face", "uint32_group_default_page", "uint32_group_info_seq", "uint32_group_roaming_time", "string_group_name", "string_group_memo", "string_group_finger_memo", "string_group_class_text", "uint32_group_alliance_code", "uint32_group_extra_adm_num", "uint64_group_uin", "uint32_group_cur_msg_seq", "uint32_group_last_msg_time", "string_group_question", "string_group_answer", "uint32_group_visitor_max_num", "uint32_group_visitor_cur_num", "uint32_level_name_seq", "uint32_group_admin_max_num", "uint32_group_aio_skin_timestamp", "uint32_group_board_skin_timestamp", "string_group_aio_skin_url", "string_group_board_skin_url", "uint32_group_cover_skin_timestamp", "string_group_cover_skin_url", "uint32_group_grade", "uint32_active_member_num", "uint32_certification_type", "string_certification_text", "string_group_rich_finger_memo", "rpt_tag_record", "group_geo_info", "uint32_head_portrait_seq", "msg_head_portrait", "uint32_shutup_timestamp", "uint32_shutup_timestamp_me", "uint32_create_source_flag", "uint32_cmduin_msg_seq", "uint32_cmduin_join_time", "uint32_cmduin_uin_flag", "uint32_cmduin_flag_ex", "uint32_cmduin_new_mobile_flag", "uint32_cmduin_read_msg_seq", "uint32_cmduin_last_msg_time", "uint32_group_type_flag", "uint32_app_privilege_flag", "st_group_ex_info", "uint32_group_sec_level", "uint32_group_sec_level_info", "uint32_cmduin_privilege", "string_poid_info", "uint32_cmduin_flag_ex2", "uint64_conf_uin", "uint32_conf_max_msg_seq", "uint32_conf_to_group_time", "uint32_password_redbag_time", "uint64_subscription_uin", "uint32_member_list_change_seq", "uint32_membercard_seq", "uint64_root_id", "uint64_parent_id", "uint32_team_seq", "uint64_history_msg_begin_time", "uint64_invite_no_auth_num_limit", "uint32_cmduin_history_msg_seq", "uint32_cmduin_join_msg_seq", "uint32_group_flagext3", "uint32_group_open_appid", "uint32_is_conf_group", "uint32_is_modify_conf_group_face", "uint32_is_modify_conf_group_name", "uint32_no_finger_open_flag", "uint32_no_code_finger_open_flag", "uint32_auto_agree_join_group_user_num_for_normal_group", "uint32_auto_agree_join_group_user_num_for_conf_group", "uint32_is_allow_conf_group_member_nick", "uint32_is_allow_conf_group_member_at_all", "uint32_is_allow_conf_group_member_modify_group_name", "string_long_group_name", "uint32_cmduin_join_real_msg_seq", "uint32_is_group_freeze", "uint32_msg_limit_frequency", "bytes_join_group_auth", "uint32_hl_guild_appid", "uint32_hl_guild_sub_type", "uint32_hl_guild_orgid", "uint32_is_allow_hl_guild_binary", "uint32_cmduin_ringtone_id", "uint32_group_flagext4", "uint32_group_freeze_reason", "uint32_is_allow_recall_msg", "uint32_important_msg_latest_seq", "bytes_group_school_info", "uint32_appeal_deadline", "st_group_card_prefix", "uint64_alliance_id", "uint32_cmduin_flag_ex3_grocery", "uint32_group_info_ext_seq", "st_group_info_ext", "bytes_cmduin_group_remark_name"}, new Object[]{0L, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, 0, 0, 0L, 0, 0, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, 0, 0, 0, 0, 0, 0, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, 0, ByteStringMicro.EMPTY, 0, 0, 0, ByteStringMicro.EMPTY, ByteStringMicro.EMPTY, null, null, 0, null, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, null, 0, 0, 0, ByteStringMicro.EMPTY, 0, 0L, 0, 0, 0, 0L, 0, 0, 0L, 0L, 0, 0L, 0L, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ByteStringMicro.EMPTY, 0, 0, 0, ByteStringMicro.EMPTY, 0, 0, 0, 0, 0, 0, 0, 0, 0, ByteStringMicro.EMPTY, 0, null, 0L, 0, 0, null, ByteStringMicro.EMPTY}, GroupInfo.class);
        public final PBBytesField bytes_cmduin_group_remark_name = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField bytes_group_school_info = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField bytes_join_group_auth = PBField.initBytes(ByteStringMicro.EMPTY);
        public GroupGeoInfo group_geo_info = new GroupGeoInfo();
        public GroupHeadPortrait msg_head_portrait = new GroupHeadPortrait();
        public final PBRepeatMessageField<TagRecord> rpt_tag_record = PBField.initRepeatMessage(TagRecord.class);
        public GroupCardPrefix st_group_card_prefix = new GroupCardPrefix();
        public GroupExInfoOnly st_group_ex_info = new GroupExInfoOnly();
        public oidb_0xef0.GroupInfoExt st_group_info_ext = new oidb_0xef0.GroupInfoExt();
        public final PBBytesField string_certification_text = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_aio_skin_url = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_answer = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_board_skin_url = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_class_text = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_cover_skin_url = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_finger_memo = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_memo = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_name = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_question = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_group_rich_finger_memo = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_long_group_name = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField string_poid_info = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBUInt32Field uint32_active_member_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_app_privilege_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_appeal_deadline = PBField.initUInt32(0);
        public final PBUInt32Field uint32_auto_agree_join_group_user_num_for_conf_group = PBField.initUInt32(0);
        public final PBUInt32Field uint32_auto_agree_join_group_user_num_for_normal_group = PBField.initUInt32(0);
        public final PBUInt32Field uint32_certification_type = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_flag_ex = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_flag_ex2 = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_flag_ex3_grocery = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_history_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_join_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_join_real_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_join_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_last_msg_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_new_mobile_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_privilege = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_read_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_ringtone_id = PBField.initUInt32(0);
        public final PBUInt32Field uint32_cmduin_uin_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_conf_max_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_conf_to_group_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_create_source_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_admin_max_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_aio_skin_timestamp = PBField.initUInt32(0);
        public final PBRepeatField<Integer> uint32_group_alliance_code = PBField.initRepeat(PBUInt32Field.__repeatHelper__);
        public final PBUInt32Field uint32_group_board_skin_timestamp = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_class_ext = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_cover_skin_timestamp = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_create_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_cur_msg_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_default_page = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_extra_adm_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_face = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_flag_ext = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_flagext3 = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_flagext4 = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_freeze_reason = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_grade = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_info_ext_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_info_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_last_msg_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_level = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_member_max_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_member_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_open_appid = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_option = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_roaming_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_sec_level = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_sec_level_info = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_special_class = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_type_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_visitor_cur_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_group_visitor_max_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_head_portrait_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_hl_guild_appid = PBField.initUInt32(0);
        public final PBUInt32Field uint32_hl_guild_orgid = PBField.initUInt32(0);
        public final PBUInt32Field uint32_hl_guild_sub_type = PBField.initUInt32(0);
        public final PBUInt32Field uint32_important_msg_latest_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_allow_conf_group_member_at_all = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_allow_conf_group_member_modify_group_name = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_allow_conf_group_member_nick = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_allow_hl_guild_binary = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_allow_recall_msg = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_conf_group = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_group_freeze = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_modify_conf_group_face = PBField.initUInt32(0);
        public final PBUInt32Field uint32_is_modify_conf_group_name = PBField.initUInt32(0);
        public final PBUInt32Field uint32_level_name_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_member_list_change_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_membercard_seq = PBField.initUInt32(0);
        public final PBUInt32Field uint32_msg_limit_frequency = PBField.initUInt32(0);
        public final PBUInt32Field uint32_no_code_finger_open_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_no_finger_open_flag = PBField.initUInt32(0);
        public final PBUInt32Field uint32_password_redbag_time = PBField.initUInt32(0);
        public final PBUInt32Field uint32_shutup_timestamp = PBField.initUInt32(0);
        public final PBUInt32Field uint32_shutup_timestamp_me = PBField.initUInt32(0);
        public final PBUInt32Field uint32_team_seq = PBField.initUInt32(0);
        public final PBUInt64Field uint64_alliance_id = PBField.initUInt64(0);
        public final PBUInt64Field uint64_conf_uin = PBField.initUInt64(0);
        public final PBUInt64Field uint64_group_owner = PBField.initUInt64(0);
        public final PBUInt64Field uint64_group_uin = PBField.initUInt64(0);
        public final PBUInt64Field uint64_history_msg_begin_time = PBField.initUInt64(0);
        public final PBUInt64Field uint64_invite_no_auth_num_limit = PBField.initUInt64(0);
        public final PBUInt64Field uint64_parent_id = PBField.initUInt64(0);
        public final PBUInt64Field uint64_root_id = PBField.initUInt64(0);
        public final PBUInt64Field uint64_subscription_uin = PBField.initUInt64(0);
    }

    /* compiled from: P */
    public final class ReqBody extends MessageMicro<ReqBody> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 18, 24}, new String[]{"uint32_appid", "stzreqgroupinfo", "uint32_pc_client_version"}, new Object[]{0, null, 0}, ReqBody.class);
        public final PBRepeatMessageField<ReqGroupInfo> stzreqgroupinfo = PBField.initRepeatMessage(ReqGroupInfo.class);
        public final PBUInt32Field uint32_appid = PBField.initUInt32(0);
        public final PBUInt32Field uint32_pc_client_version = PBField.initUInt32(0);
    }

    /* compiled from: P */
    public final class ReqGroupInfo extends MessageMicro<ReqGroupInfo> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 18, 24}, new String[]{"uint64_group_code", "stgroupinfo", "uint32_last_get_group_name_time"}, new Object[]{0L, null, 0}, ReqGroupInfo.class);
        public GroupInfo stgroupinfo = new GroupInfo();
        public final PBUInt32Field uint32_last_get_group_name_time = PBField.initUInt32(0);
        public final PBUInt64Field uint64_group_code = PBField.initUInt64(0);
    }

    /* compiled from: P */
    public final class RspBody extends MessageMicro<RspBody> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{10, 18}, new String[]{"stzrspgroupinfo", "str_errorinfo"}, new Object[]{null, ByteStringMicro.EMPTY}, RspBody.class);
        public final PBBytesField str_errorinfo = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBRepeatMessageField<RspGroupInfo> stzrspgroupinfo = PBField.initRepeatMessage(RspGroupInfo.class);
    }

    /* compiled from: P */
    public final class RspGroupInfo extends MessageMicro<RspGroupInfo> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16, 26}, new String[]{"uint64_group_code", "uint32_result", "stgroupinfo"}, new Object[]{0L, 0, null}, RspGroupInfo.class);
        public GroupInfo stgroupinfo = new GroupInfo();
        public final PBUInt32Field uint32_result = PBField.initUInt32(0);
        public final PBUInt64Field uint64_group_code = PBField.initUInt64(0);
    }

    /* compiled from: P */
    public final class TagRecord extends MessageMicro<TagRecord> {
        static final MessageMicro.FieldMap __fieldMap__ = MessageMicro.initFieldMap(new int[]{8, 16, 26, 32, 40, 48, 56, 66}, new String[]{"uint64_from_uin", "uint64_group_code", "bytes_tag_id", "uint64_set_time", "uint32_good_num", "uint32_bad_num", "uint32_tag_len", "bytes_tag_value"}, new Object[]{0L, 0L, ByteStringMicro.EMPTY, 0L, 0, 0, 0, ByteStringMicro.EMPTY}, TagRecord.class);
        public final PBBytesField bytes_tag_id = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBBytesField bytes_tag_value = PBField.initBytes(ByteStringMicro.EMPTY);
        public final PBUInt32Field uint32_bad_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_good_num = PBField.initUInt32(0);
        public final PBUInt32Field uint32_tag_len = PBField.initUInt32(0);
        public final PBUInt64Field uint64_from_uin = PBField.initUInt64(0);
        public final PBUInt64Field uint64_group_code = PBField.initUInt64(0);
        public final PBUInt64Field uint64_set_time = PBField.initUInt64(0);
    }

    private oidb_0x88d() {
    }
}