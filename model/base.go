package model

func Tables() []interface{} {
	return []interface{}{
		new(Option),
		new(User),
		new(UserProfile),
		new(UserInvitation),
	}
}
