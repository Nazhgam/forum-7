package repo

import "forum/entity"

type IProfile interface {
	CreateProfile(p *entity.Profile) error
	GetProfileByID(id int64) (*entity.Profile, error)
	UpdateProfile(p entity.Profile) error
	DeleteProfileByID(id int64) error
}

func (r repo) CreateProfile(p *entity.Profile) error {
	stmt, err := r.db.Prepare("INSERT INTO profiles (user_id, name, bio, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the profile table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.UserId, p.Name, p.Bio, p.ImageUrl, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into profile table: %s\n", err.Error())
		return err
	}

	return nil
}

// getProfileByID retrieves a profile from the Profile table by ID
func (r repo) GetProfileByID(id int64) (*entity.Profile, error) {
	stmt, err := r.db.Prepare("SELECT id, user_id, name, bio, image_url, created_at, updated_at FROM profiles WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get profile by id from profile table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var profile entity.Profile
	err = stmt.QueryRow(id).Scan(&profile.Id, &profile.UserId, &profile.Name, &profile.Bio, &profile.ImageUrl, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		r.log.Printf("error while to query row and scan profile to get by id: %s\n", err.Error())
		return nil, err
	}

	return &profile, nil
}

// updateProfile updates an existing profile in the Profile table
func (r repo) UpdateProfile(p entity.Profile) error {
	stmt, err := r.db.Prepare("UPDATE profiles SET user_id = ?, name = ?, bio = ?, image_url = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in profile table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.UserId, p.Name, p.Bio, p.ImageUrl, p.UpdatedAt, p.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in profile table: %s\n", err.Error())
		return err
	}

	return nil
}

// deleteProfileByID deletes a profile from the Profile table by ID
func (r repo) DeleteProfileByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM profiles WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete profile by id in profile table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete profile by id in profile table: %s\n", err.Error())
		return err
	}

	return nil
}
