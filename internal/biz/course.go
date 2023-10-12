package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "kubecit-service/api/helloworld/v1"
)

// Category is a Category model.
type Category struct {
	CategoryName string
	Id           int32
	ParentId     int32
	Level        int
	Children     []*Category
}

type Course struct {
	Id            int
	Level         int32
	Name          string
	Detail        string
	Cover         string
	Price         int32
	Tags          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        int32
	CategoryId    int
	People        int32
	Duration      int32
	Score         int32
	TeacherId     int
	TeacherAvatar string
	TeacherLevel  int32
}

type CourseChapterLessons struct {
	Course
	Chapters ChapterLessons
}

type Chapter struct {
	Id           int
	Name         string
	ReleasedTime time.Time
	Description  string
	Sort         int
	CourseId     int
}

type Lesson struct {
	Id            int
	Name          string
	ReleasedTime  time.Time
	Sort          int
	Type          int
	StoragePath   string
	Source        string
	Courseware    string
	IsFreePreview int
	ChapterId     int
}

type ChapterLessons struct {
	*Chapter
	Lessons []*Lesson
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	ListAll(ctx context.Context) ([]*Category, error)
	ListByLevelAndCategory(ctx context.Context, level, category *int32) ([]*Category, error)
	Create(context.Context, *Category) error
	ListSubCategories(ctx context.Context, parentID int32) ([]*Category, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, id int, name string) error
}

// CourseRepo is a Course repo.
type CourseRepo interface {
	SearchCourse(ctx context.Context, pageNum, pageSize *int32, categoryIds []int, level pb.CourseLevel, order *int32, name *string) ([]*Course, int32, error)
	UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error)
	ReviewCourse(ctx context.Context, id int, status int32) (*Course, error)
	CreateCourse(ctx context.Context, course *Course) (*Course, error)
	GetCourse(ctx context.Context, id int) (*Course, error)
	DeleteCourse(ctx context.Context, id int) (int, error)

	CreateChapter(ctx context.Context, chapter *Chapter) (*Chapter, error)
	DeleteChapter(ctx context.Context, id int) (int, error)
	ListChapters(ctx context.Context, courseId int) ([]*Chapter, error)
	UpdateChapter(ctx context.Context, id int, chapter *Chapter) (*Chapter, error)

	GetLesson(ctx context.Context, id int) (*Lesson, error)
	CreateLesson(ctx context.Context, lesson *Lesson) (*Lesson, error)
	DeleteLesson(ctx context.Context, id int) (int, error)
	ListLessons(ctx context.Context, chapterId int) ([]*Lesson, error)
	UpdateLesson(ctx context.Context, id int, chapter *Lesson) (*Lesson, error)
}

// CourseUsecase is a Category usecase.
type CourseUsecase struct {
	repo       CategoryRepo
	courseRepo CourseRepo

	log *log.Helper
}

// NewCourseUsecase new a Category usecase.
func NewCourseUsecase(repo CategoryRepo, courseRepo CourseRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{
		repo:       repo,
		courseRepo: courseRepo,
		log:        log.NewHelper(logger),
	}
}

func (uc *CourseUsecase) ListCategory(ctx context.Context, level, categoryId *int32) ([]*Category, error) {
	return uc.repo.ListByLevelAndCategory(ctx, level, categoryId)
}

func (uc *CourseUsecase) ListCategoryV2(ctx context.Context, level, categoryId *int32) ([]*Category, error) {
	categorires, err := uc.repo.ListByLevelAndCategory(ctx, level, categoryId)
	if err != nil {
		return nil, err
	}

	for _, v := range categorires {
		sub, err := uc.repo.ListSubCategories(ctx, v.Id)
		if err != nil {
			continue
		}
		v.Children = sub
	}
	return categorires, nil
}

func (uc *CourseUsecase) CreateCategory(ctx context.Context, category *Category) error {
	return uc.repo.Create(ctx, category)
}

func (uc *CourseUsecase) DeleteCategory(ctx context.Context, id int32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *CourseUsecase) UpdateCategory(ctx context.Context, id int, name string) error {
	return uc.repo.Update(ctx, id, name)
}

type SearchFilterParam struct {
	PageNum          *int32
	PageSize         *int32
	SecondCategoryId *int32
	FirstCategoryId  *int32
	Level            pb.CourseLevel
	Order            *int32
	Name             *string
}

func (uc *CourseUsecase) SearchCourse(ctx context.Context, filter *SearchFilterParam) ([]*Course, int32, error) {

	var categoryIds []int
	if filter.SecondCategoryId == nil {
		if filter.FirstCategoryId != nil {
			subCategories, err := uc.repo.ListSubCategories(ctx, *filter.FirstCategoryId)
			if err != nil {
				return nil, 0, err
			}
			for _, v := range subCategories {
				categoryIds = append(categoryIds, int(v.Id))
			}
		}

	} else {
		categoryIds = append(categoryIds, int(*filter.SecondCategoryId))
	}

	return uc.courseRepo.SearchCourse(ctx, filter.PageNum, filter.PageSize, categoryIds, filter.Level, filter.Order, filter.Name)
}

func (uc *CourseUsecase) UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error) {
	return uc.courseRepo.UpdateCourse(ctx, id, course)
}

func (uc *CourseUsecase) ReviewCourse(ctx context.Context, id int, status int32) (*Course, error) {
	return uc.courseRepo.ReviewCourse(ctx, id, status)
}

func (uc *CourseUsecase) CreateCourse(ctx context.Context, course *Course) (*Course, error) {
	return uc.courseRepo.CreateCourse(ctx, course)
}

func (uc *CourseUsecase) GetCourse(ctx context.Context, id int) (*Course, error) {
	return uc.courseRepo.GetCourse(ctx, id)
}

func (uc *CourseUsecase) DeleteCourse(ctx context.Context, id int) (int, error) {
	return uc.courseRepo.DeleteCourse(ctx, id)
}

func (uc *CourseUsecase) CreateChapter(ctx context.Context, chapter *Chapter) (*Chapter, error) {
	return uc.courseRepo.CreateChapter(ctx, chapter)
}

func (uc *CourseUsecase) DeleteChapter(ctx context.Context, id int) (int, error) {
	return uc.courseRepo.DeleteChapter(ctx, id)
}

func (uc *CourseUsecase) ListChapters(ctx context.Context, courseId int) ([]*Chapter, error) {
	return uc.courseRepo.ListChapters(ctx, courseId)
}

func (uc *CourseUsecase) UpdateChapter(ctx context.Context, id int, chapter *Chapter) (*Chapter, error) {
	return uc.courseRepo.UpdateChapter(ctx, id, chapter)
}

func (uc *CourseUsecase) CreateLesson(ctx context.Context, lesson *Lesson) (*Lesson, error) {
	return uc.courseRepo.CreateLesson(ctx, lesson)
}
func (uc *CourseUsecase) DeleteLesson(ctx context.Context, id int) (int, error) {
	return uc.courseRepo.DeleteLesson(ctx, id)
}

func (uc *CourseUsecase) ListChapterLessonsWithCourseId(ctx context.Context, courseId int) ([]*ChapterLessons, error) {
	chapters, err := uc.courseRepo.ListChapters(ctx, courseId)
	if err != nil || len(chapters) == 0 {
		return nil, err
	}
	result := make([]*ChapterLessons, 0, len(chapters))
	for _, chapter := range chapters {
		lessons, err := uc.courseRepo.ListLessons(ctx, chapter.Id)
		if err != nil {
			return nil, err
		}
		result = append(result, &ChapterLessons{
			Chapter: chapter,
			Lessons: lessons,
		})
	}
	return result, nil
}

func (uc *CourseUsecase) UpdateLesson(ctx context.Context, id int, chapter *Lesson) (*Lesson, error) {
	return uc.courseRepo.UpdateLesson(ctx, id, chapter)
}
