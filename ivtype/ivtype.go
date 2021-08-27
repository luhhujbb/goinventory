package ivtype

type Store struct {
    Key string
    Bucket string
    Type string
    Atomic bool // Indicate what is send
    Format string // Only for nonatomic store yaml or json
}

type Update struct {
    Id string //only for atomic update
    Body string
}

type ApiResponse struct {
    State string "json: state"
    Data interface{} "json: data"
}
