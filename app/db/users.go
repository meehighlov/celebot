package db

import (
	"fmt"
	"log"
)

func (user *User) Save() error {
	stmt, err := Client.Prepare(
        "INSERT INTO user(id, name, tgusername, chatid, birthday) " +
        "VALUES($1, $2, $3, $4, $5) " +
        "ON CONFLICT(id) DO UPDATE SET name=$2, tgusername=$3, chatid=$4, birthday=$5 " +
        "RETURNING id;",
    )
    if err != nil {
        log.Println("Error when trying to prepare statement for saving user: " + err.Error())
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(&user.ID, &user.Name, &user.TGusername, &user.ChatId, &user.Birthday).Scan(&user.ID)
    if insertErr != nil {
        log.Println("Error when trying to save user: " + err.Error())
        return err
    }
    log.Println("User created/updated")

    return nil
}

func (user *User) Get() error {
    stmt, err := Client.Prepare("SELECT id, name, tgusername, chatid, birthday FROM user WHERE id=$1;")
    if err != nil {
        log.Println("Error when trying to prepare statement for getting user by id: " + err.Error())
        return err
    }
    defer stmt.Close()

    result := stmt.QueryRow(&user.ID)
    if err := result.Scan(&user.ID, &user.Name, &user.TGusername, &user.ChatId, &user.Birthday); err != nil {
        log.Println("Error when trying to get User by ID: " + err.Error())
        return err
    }

    return nil
}

func (user *User) GetById(fetchFriends bool) error {
	stmt, err := Client.Prepare("SELECT id, name, tgusername, chatid, birthday FROM user WHERE id=$1;")
    if err != nil {
        log.Println("Error when trying to prepare statement for getting user by id")
        return err
    }
    defer stmt.Close()

    result := stmt.QueryRow(&user.ID)

	if err := result.Scan(&user.ID, &user.Name, &user.TGusername, &user.ChatId, &user.Birthday); err != nil {
        log.Println("Error when trying to get User by ID")
        return err
    }

	if fetchFriends {
		stmt, err := Client.Prepare("SELECT id, name, birthday, userid, chatid FROM friend WHERE userid=$1;")
		if err != nil {
			log.Println("Error when trying to prepare statement for fetching friends for user")
			return err
		}

		results, err := stmt.Query(&user.ID)
		if err != nil {
			log.Println("Error when fetching friends for user")
			return err
		}

		for results.Next() {
			friend := Friend{}
			err := results.Scan(&friend.ID, &friend.Name, &friend.BirthDay, &friend.UserId, &friend.ChatId)
			if err != nil {
				log.Println("Error when fetching friends for user")
				continue
			}
			user.Friends = append(user.Friends, friend)
		}
	}

    return nil
}

func (user *User) GetFriendsByBirthDate(birthDay string, limit, offset int) error {
    stmt, err := Client.Prepare(
        "SELECT id, name, birthday, userid, chatid FROM friend WHERE birthday LIKE $1 OR birthday LIKE $2 LIMIT $3 OFFSET $4;",
    )
    if err != nil {
        log.Println("Error when trying to prepare statement for fetching friends for user")
        return err
    }

    results, err := stmt.Query(birthDay + ".%", birthDay, limit, offset)
    if err != nil {
        log.Println("Error when fetching friends for user by birthday")
        return err
    }

    for results.Next() {
        friend := Friend{}
        err := results.Scan(&friend.ID, &friend.Name, &friend.BirthDay, &friend.UserId, &friend.ChatId)
        if err != nil {
            log.Println("Error when fetching friends for user by birthday")
            continue
        }
        user.Friends = append(user.Friends, friend)
    }

    return nil
}

func (friend *Friend) Save() error {
	stmt, err := Client.Prepare("INSERT INTO friend(name, birthday, userid, chatid) VALUES($1, $2, $3, $4) RETURNING id;")
    if err != nil {
        log.Println("Error when trying to prepare statement: " + err.Error())
        return err
    }
    defer stmt.Close()

    insertErr := stmt.QueryRow(friend.Name, friend.BirthDay, friend.UserId, friend.ChatId).Scan(&friend.ID)
    if insertErr != nil {
        _ = fmt.Sprintf("Error when trying to save friend with id: %d", friend.ID)
        return err
    }
    log.Println("Friend added")

    return nil
}
