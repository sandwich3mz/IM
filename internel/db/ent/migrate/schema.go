// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TFriendColumns holds the columns for the "t_friend" table.
	TFriendColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "relationship", Type: field.TypeInt8, Comment: "用户关系"},
		{Name: "remark", Type: field.TypeString, Comment: "好友备注"},
		{Name: "face_url", Type: field.TypeString, Comment: "头像"},
		{Name: "nickname", Type: field.TypeString, Comment: "昵称"},
		{Name: "owner_user_id", Type: field.TypeInt64, Comment: "当前登录用户 ID"},
		{Name: "friend_user_id", Type: field.TypeInt64, Comment: "好友 ID"},
	}
	// TFriendTable holds the schema information for the "t_friend" table.
	TFriendTable = &schema.Table{
		Name:       "t_friend",
		Columns:    TFriendColumns,
		PrimaryKey: []*schema.Column{TFriendColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "t_friend_t_user_owner_user_friend",
				Columns:    []*schema.Column{TFriendColumns[8]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "t_friend_t_user_friend_user_friend",
				Columns:    []*schema.Column{TFriendColumns[9]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TGroupColumns holds the columns for the "t_group" table.
	TGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "group_name", Type: field.TypeString, Comment: "群聊名称"},
		{Name: "notification", Type: field.TypeString, Comment: "群公告"},
		{Name: "relationship", Type: field.TypeInt8, Comment: "用户关系"},
		{Name: "face_url", Type: field.TypeString, Comment: "头像"},
		{Name: "count", Type: field.TypeInt64, Comment: "群员数量"},
		{Name: "owner_user_id", Type: field.TypeInt64, Comment: "群主 ID"},
	}
	// TGroupTable holds the schema information for the "t_group" table.
	TGroupTable = &schema.Table{
		Name:       "t_group",
		Columns:    TGroupColumns,
		PrimaryKey: []*schema.Column{TGroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "t_group_t_user_user_group",
				Columns:    []*schema.Column{TGroupColumns[9]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TGroupMemberColumns holds the columns for the "t_group_member" table.
	TGroupMemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "nickname", Type: field.TypeInt8, Comment: "群成员昵称"},
		{Name: "face_url", Type: field.TypeString, Comment: "头像"},
		{Name: "count", Type: field.TypeInt64, Comment: "群员数量"},
		{Name: "remark", Type: field.TypeString, Comment: "备注"},
		{Name: "group_id", Type: field.TypeInt64, Comment: "群 ID"},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户 ID"},
	}
	// TGroupMemberTable holds the schema information for the "t_group_member" table.
	TGroupMemberTable = &schema.Table{
		Name:       "t_group_member",
		Columns:    TGroupMemberColumns,
		PrimaryKey: []*schema.Column{TGroupMemberColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "t_group_member_t_group_group_member",
				Columns:    []*schema.Column{TGroupMemberColumns[8]},
				RefColumns: []*schema.Column{TGroupColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "t_group_member_t_user_user_group_member",
				Columns:    []*schema.Column{TGroupMemberColumns[9]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TMessageColumns holds the columns for the "t_message" table.
	TMessageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "send_at", Type: field.TypeTime, Comment: "发送时间"},
		{Name: "session_type", Type: field.TypeEnum, Comment: "会话类型", Enums: []string{"unknown", "single", "group"}, Default: "unknown"},
		{Name: "content_type", Type: field.TypeEnum, Comment: "消息类型", Enums: []string{"unknown", "textMessage"}, Default: "unknown"},
		{Name: "seq", Type: field.TypeInt32, Comment: "消息序列号", Default: -1},
		{Name: "status", Type: field.TypeEnum, Comment: "消息状态", Enums: []string{"unknown", "sending", "succeed", "failed"}, Default: "unknown"},
		{Name: "text_elem", Type: field.TypeString, Size: 2147483647, Comment: "文本信息"},
		{Name: "send_id", Type: field.TypeInt64, Comment: "发送者 ID"},
		{Name: "receive_id", Type: field.TypeInt64, Comment: "发送者 ID"},
	}
	// TMessageTable holds the schema information for the "t_message" table.
	TMessageTable = &schema.Table{
		Name:       "t_message",
		Columns:    TMessageColumns,
		PrimaryKey: []*schema.Column{TMessageColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "t_message_t_user_send_msg",
				Columns:    []*schema.Column{TMessageColumns[10]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "t_message_t_user_receive_msg",
				Columns:    []*schema.Column{TMessageColumns[11]},
				RefColumns: []*schema.Column{TUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TUserColumns holds the columns for the "t_user" table.
	TUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "nick_name", Type: field.TypeString, Comment: "用户昵称", Default: "请设置昵称"},
		{Name: "email", Type: field.TypeString, Comment: "邮箱", Default: ""},
		{Name: "password", Type: field.TypeString, Comment: "密码", Default: ""},
		{Name: "status", Type: field.TypeEnum, Comment: "用户状态", Enums: []string{"online", "offline"}, Default: "offline"},
		{Name: "last_online_at", Type: field.TypeTime, Comment: "最后在线时间"},
	}
	// TUserTable holds the schema information for the "t_user" table.
	TUserTable = &schema.Table{
		Name:       "t_user",
		Columns:    TUserColumns,
		PrimaryKey: []*schema.Column{TUserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TFriendTable,
		TGroupTable,
		TGroupMemberTable,
		TMessageTable,
		TUserTable,
	}
)

func init() {
	TFriendTable.ForeignKeys[0].RefTable = TUserTable
	TFriendTable.ForeignKeys[1].RefTable = TUserTable
	TFriendTable.Annotation = &entsql.Annotation{
		Table: "t_friend",
	}
	TGroupTable.ForeignKeys[0].RefTable = TUserTable
	TGroupTable.Annotation = &entsql.Annotation{
		Table: "t_group",
	}
	TGroupMemberTable.ForeignKeys[0].RefTable = TGroupTable
	TGroupMemberTable.ForeignKeys[1].RefTable = TUserTable
	TGroupMemberTable.Annotation = &entsql.Annotation{
		Table: "t_group_member",
	}
	TMessageTable.ForeignKeys[0].RefTable = TUserTable
	TMessageTable.ForeignKeys[1].RefTable = TUserTable
	TMessageTable.Annotation = &entsql.Annotation{
		Table: "t_message",
	}
	TUserTable.Annotation = &entsql.Annotation{
		Table: "t_user",
	}
}
