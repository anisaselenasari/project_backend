// CREATE TABLE public.artikel (
// 	id uuid NOT NULL DEFAULT uuid_generate_v4(),
// 	title varchar(150) NOT NULL,
// 	kontent text NOT NULL,
// 	created_at timestamp NOT NULL,
// 	updated_at timestamptz NULL,
// 	deleted_at timestamp NULL,
// // )

// CREATE TABLE public.komentar (
// 	id uuid NOT NULL DEFAULT uuid_generate_v4(),
// 	artikel_id uuid NOT NULL,
// 	komentar text NOT NULL,
// 	created_at timestamp NOT NULL,
// 	updated_at timestamptz NULL,
// 	deleted_at timestamp NULL
// );

// type Artikel struct {
// 	ID        uuid.UUID  `gorm:"column:id" json:"artikel_id"`
// 	Title     string     `gorm:"column:title" json:"title"`
// 	Kontent   string     `gorm:"column:kontent" json:"kontent"`
// 	Komentar  []Komentar `gorm:"Foreignkey:Artikel_ID;association_foreignkey:ID;" json:"komentar"`
// 	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
// }

// type Komentar struct {
// 	ID         uuid.UUID  `gorm:"column:id" json:"komentar_id"`
// 	Artikel_ID string     `gorm:"column:artikel_id" json:"artikel_id"`
// 	Komentar   string     `gorm:"column:komentar" json:"komentar"`
// 	CreatedAt  time.Time  `gorm:"column:created_at" json:"-"`
// 	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"-"`
// 	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"-"`
// }