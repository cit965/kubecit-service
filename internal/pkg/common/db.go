package common

func ConvertPageSize(pageNum, pageSize *int32) (limit, offset int) {

	if pageNum == nil || *pageNum <= 0 {
		*pageNum = 1
	}
	if pageSize == nil || *pageSize <= 0 {
		*pageSize = 20
	}
	limit = int(*pageSize)
	offset = int((*pageNum - 1) * (*pageSize))
	return
}
