package fakedb

var (
	nextPostID int64 = 0
)

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type PostRepository struct{}

func GeneratePostID() int64 {
	nextPostID++
	return nextPostID
}

func (p *PostRepository) Get(ID int64) *Post {
	for _, post := range Posts {
		if post.ID == ID {
			return post
		}
	}
	return nil
}

func (p *PostRepository) All() []*Post {
	return Posts
}

func (p *PostRepository) Save(post *Post) {
	Posts = append(Posts, post)
}

func (p *PostRepository) indexOf(ID int64) int {
	for index, post := range Posts {
		if post.ID == ID {
			return index
		}
	}
	return -1
}

func (p *PostRepository) ByUserID(ID int64) (result []*Post) {
	for _, post := range Posts {
		if post.User != nil && post.User.ID == ID {
			result = append(result, post)
		}
	}
	return
}

func (p *PostRepository) Delete(ID int64) error {
	index := p.indexOf(ID)
	if index == -1 {
		return ErrNotFound
	}
	Posts = append(Posts[:index], Posts[index+1:]...)
	return nil
}
