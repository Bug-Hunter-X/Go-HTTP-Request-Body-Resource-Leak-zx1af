func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    defer r.Body.Close() // Crucial for proper resource management
    // ... rest of the request handling ...
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    //Process the body
}