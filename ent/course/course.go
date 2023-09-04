// Code generated by ent, DO NOT EDIT.

package course

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsRecommend holds the string denoting the isrecommend field in the database.
	FieldIsRecommend = "is_recommend"
	// FieldIsIntegral holds the string denoting the isintegral field in the database.
	FieldIsIntegral = "is_integral"
	// FieldSecondCategory holds the string denoting the secondcategory field in the database.
	FieldSecondCategory = "second_category"
	// FieldSaleType holds the string denoting the saletype field in the database.
	FieldSaleType = "sale_type"
	// FieldDiscountPrice holds the string denoting the discountprice field in the database.
	FieldDiscountPrice = "discount_price"
	// FieldFirstCategoryName holds the string denoting the firstcategoryname field in the database.
	FieldFirstCategoryName = "first_category_name"
	// FieldTeachingType holds the string denoting the teachingtype field in the database.
	FieldTeachingType = "teaching_type"
	// FieldCourseLevel holds the string denoting the courselevel field in the database.
	FieldCourseLevel = "course_level"
	// FieldUpdateBy holds the string denoting the updateby field in the database.
	FieldUpdateBy = "update_by"
	// FieldLecturerName holds the string denoting the lecturername field in the database.
	FieldLecturerName = "lecturer_name"
	// FieldPurchaseCnt holds the string denoting the purchasecnt field in the database.
	FieldPurchaseCnt = "purchase_cnt"
	// FieldTotalHour holds the string denoting the totalhour field in the database.
	FieldTotalHour = "total_hour"
	// FieldBizCourseDetail holds the string denoting the bizcoursedetail field in the database.
	FieldBizCourseDetail = "biz_course_detail"
	// FieldCourseCover holds the string denoting the coursecover field in the database.
	FieldCourseCover = "course_cover"
	// FieldExt3 holds the string denoting the ext3 field in the database.
	FieldExt3 = "ext3"
	// FieldExt2 holds the string denoting the ext2 field in the database.
	FieldExt2 = "ext2"
	// FieldBizCourseChapters holds the string denoting the bizcoursechapters field in the database.
	FieldBizCourseChapters = "biz_course_chapters"
	// FieldExt1 holds the string denoting the ext1 field in the database.
	FieldExt1 = "ext1"
	// FieldSalePrice holds the string denoting the saleprice field in the database.
	FieldSalePrice = "sale_price"
	// FieldBizCourseTeacher holds the string denoting the bizcourseteacher field in the database.
	FieldBizCourseTeacher = "biz_course_teacher"
	// FieldBizCourseAttachments holds the string denoting the bizcourseattachments field in the database.
	FieldBizCourseAttachments = "biz_course_attachments"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCourseName holds the string denoting the coursename field in the database.
	FieldCourseName = "course_name"
	// FieldCreateBy holds the string denoting the createby field in the database.
	FieldCreateBy = "create_by"
	// FieldPurchaseCounter holds the string denoting the purchasecounter field in the database.
	FieldPurchaseCounter = "purchase_counter"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldClicks holds the string denoting the clicks field in the database.
	FieldClicks = "clicks"
	// FieldSecondCategoryName holds the string denoting the secondcategoryname field in the database.
	FieldSecondCategoryName = "second_category_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the course in the database.
	Table = "courses"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldIsRecommend,
	FieldIsIntegral,
	FieldSecondCategory,
	FieldSaleType,
	FieldDiscountPrice,
	FieldFirstCategoryName,
	FieldTeachingType,
	FieldCourseLevel,
	FieldUpdateBy,
	FieldLecturerName,
	FieldPurchaseCnt,
	FieldTotalHour,
	FieldBizCourseDetail,
	FieldCourseCover,
	FieldExt3,
	FieldExt2,
	FieldBizCourseChapters,
	FieldExt1,
	FieldSalePrice,
	FieldBizCourseTeacher,
	FieldBizCourseAttachments,
	FieldUpdateTime,
	FieldTags,
	FieldCourseName,
	FieldCreateBy,
	FieldPurchaseCounter,
	FieldCreateTime,
	FieldClicks,
	FieldSecondCategoryName,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsRecommend holds the default value on creation for the "isRecommend" field.
	DefaultIsRecommend bool
	// DefaultIsIntegral holds the default value on creation for the "isIntegral" field.
	DefaultIsIntegral bool
)

// OrderOption defines the ordering options for the Course queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsRecommend orders the results by the isRecommend field.
func ByIsRecommend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRecommend, opts...).ToFunc()
}

// ByIsIntegral orders the results by the isIntegral field.
func ByIsIntegral(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsIntegral, opts...).ToFunc()
}

// BySecondCategory orders the results by the secondCategory field.
func BySecondCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecondCategory, opts...).ToFunc()
}

// BySaleType orders the results by the saleType field.
func BySaleType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSaleType, opts...).ToFunc()
}

// ByDiscountPrice orders the results by the discountPrice field.
func ByDiscountPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscountPrice, opts...).ToFunc()
}

// ByFirstCategoryName orders the results by the firstCategoryName field.
func ByFirstCategoryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstCategoryName, opts...).ToFunc()
}

// ByTeachingType orders the results by the teachingType field.
func ByTeachingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeachingType, opts...).ToFunc()
}

// ByCourseLevel orders the results by the courseLevel field.
func ByCourseLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseLevel, opts...).ToFunc()
}

// ByUpdateBy orders the results by the updateBy field.
func ByUpdateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateBy, opts...).ToFunc()
}

// ByLecturerName orders the results by the lecturerName field.
func ByLecturerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLecturerName, opts...).ToFunc()
}

// ByPurchaseCnt orders the results by the purchaseCnt field.
func ByPurchaseCnt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPurchaseCnt, opts...).ToFunc()
}

// ByTotalHour orders the results by the totalHour field.
func ByTotalHour(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalHour, opts...).ToFunc()
}

// ByBizCourseDetail orders the results by the bizCourseDetail field.
func ByBizCourseDetail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBizCourseDetail, opts...).ToFunc()
}

// ByCourseCover orders the results by the courseCover field.
func ByCourseCover(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseCover, opts...).ToFunc()
}

// ByExt3 orders the results by the ext3 field.
func ByExt3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExt3, opts...).ToFunc()
}

// ByExt2 orders the results by the ext2 field.
func ByExt2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExt2, opts...).ToFunc()
}

// ByBizCourseChapters orders the results by the bizCourseChapters field.
func ByBizCourseChapters(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBizCourseChapters, opts...).ToFunc()
}

// ByExt1 orders the results by the ext1 field.
func ByExt1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExt1, opts...).ToFunc()
}

// BySalePrice orders the results by the salePrice field.
func BySalePrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalePrice, opts...).ToFunc()
}

// ByBizCourseTeacher orders the results by the bizCourseTeacher field.
func ByBizCourseTeacher(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBizCourseTeacher, opts...).ToFunc()
}

// ByBizCourseAttachments orders the results by the bizCourseAttachments field.
func ByBizCourseAttachments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBizCourseAttachments, opts...).ToFunc()
}

// ByUpdateTime orders the results by the updateTime field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByTags orders the results by the tags field.
func ByTags(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTags, opts...).ToFunc()
}

// ByCourseName orders the results by the courseName field.
func ByCourseName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseName, opts...).ToFunc()
}

// ByCreateBy orders the results by the createBy field.
func ByCreateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateBy, opts...).ToFunc()
}

// ByPurchaseCounter orders the results by the purchaseCounter field.
func ByPurchaseCounter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPurchaseCounter, opts...).ToFunc()
}

// ByCreateTime orders the results by the createTime field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByClicks orders the results by the clicks field.
func ByClicks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClicks, opts...).ToFunc()
}

// BySecondCategoryName orders the results by the secondCategoryName field.
func BySecondCategoryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecondCategoryName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}