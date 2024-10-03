{{define "cache"}}
    func (r *UserRepo) findBySlug(slug string) User {
    var pk string
    var user User
    err := cache.GetOnce().Get(SlugCacheKey).String(&pk)
    if err == nil {
    err = cache.GetOnce().Get(fmt.Sprintf(UserCacheKey, pk)).Json(&user)
    if err == nil {
    return user
    }
    }
    r.db.Where("slug = ?", slug).First(&user)

    cache.GetOnce().Set(fmt.Sprintf(SlugCacheKey, formatInterface()), formatInterface(user.Id), 86400)
    cache.GetOnce().Set(fmt.Sprintf(SlugCacheKey, formatInterface(user.Id)), user, 86400)
    return user
    }
    func (r *UserRepo) findBySlugArr(slugArr []string) []User {
    var keys []string
    for _, v := range slugArr {
    keys = append(keys, fmt.Sprintf(UserCacheKey, formatInterface(v)))
    }
    var pkKeys []string
    multiRes, err2 := cache.GetOnce().MultiGet(keys)
    if err2 == nil {
    for _, v := range multiRes {
    var tempKey string
    err := v.String(&tempKey)
    if err == nil {
    pkKeys = append(pkKeys, tempKey)
    }
    }
    }
    if len(pkKeys) == len(slugArr) {
    get, err := cache.GetOnce().MultiGet(pkKeys)
    if err != nil {
    return nil
    }
    }
    var pk string
    var user User
    err := cache.GetOnce().Get(SlugCacheKey).String(&pk)
    if err == nil {
    err = cache.GetOnce().Get(fmt.Sprintf(UserCacheKey, pk)).Json(&user)
    if err == nil {
    return user
    }
    }
    r.db.Where("slug = ?", slug).First(&user)

    cache.GetOnce().Set(fmt.Sprintf(SlugCacheKey, formatInterface(user.Slug)), formatInterface(user.Id), 86400)
    cache.GetOnce().Set(fmt.Sprintf(SlugCacheKey, formatInterface(user.Id)), user, 86400)
    return user
    }

{{-end}}