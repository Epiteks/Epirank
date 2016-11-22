package urls

// EpitechIntranet url
const EpitechIntranet = "https://intra.epitech.eu"

// UserList url to get list of users
// Need to be formated with :
// Intranet URL
// Location  (STG, ...)
// Promo (tek1, ...)
// Offset
const UserList = "%v/user/filter/user?format=json&location=FR/%v&year=2016&active=true&promo=%v&offset=%v"
