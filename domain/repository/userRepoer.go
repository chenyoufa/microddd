package repository

type UserRepoer interface {
	Add() string
	Del() string
	Update() string
	Query() string
}
