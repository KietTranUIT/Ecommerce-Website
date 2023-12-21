function GetNewCategory() {
    const html = `<section id="section-content">
    <div class="section-sub">
        <div class="section-title">
            <h4>Dashboard/ Insert Category</h4>
        </div>
        <div class="section-view">
            <div class="section-view-header">
                <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20" viewBox="0 0 512 512">
                    <path d="M345 39.1L472.8 168.4c52.4 53 52.4 138.2 0 191.2L360.8 472.9c-9.3 9.4-24.5 9.5-33.9 .2s-9.5-24.5-.2-33.9L438.6 325.9c33.9-34.3 33.9-89.4 0-123.7L310.9 72.9c-9.3-9.4-9.2-24.6 .2-33.9s24.6-9.2 33.9 .2zM0 229.5V80C0 53.5 21.5 32 48 32H197.5c17 0 33.3 6.7 45.3 18.7l168 168c25 25 25 65.5 0 90.5L277.3 442.7c-25 25-65.5 25-90.5 0l-168-168C6.7 262.7 0 246.5 0 229.5zM144 144a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/>
                </svg>
                <h4>Insert Category</h4>
            </div>
            <div class="section-view-content">
                <div class="section-view-edit">
                    
                    <div class="section-form">
                        <form id="form-insert">
                            <ul>
                                <li>
                                    <label for="id">Id</label>
                                    <input type="text" id="id" name="id" readonly>
                                </li>
                                <li style="position:relative">
                                    <label for="name">Category Name</label>
                                    <input type="text" id="name" name="name">
                                    <svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 512 512" style="position:absolute; right:0">
                                        <path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"/>
                                    </svg>
                                </li>
                                <li>
                                    <label for="person">Group</label>
                                    <select name="person" id="person">
                                        <option value="men">Men</option>
                                        <option value="women">Women</option>
                                        <option value="kids">Kids</option>
                                    </select>
                                </li>
                                <li style="position:relative">
                                    <label for="description">Description</label>
                                    <textarea type="text" id="description" rows="7" name="description"></textarea>
                                    <svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 512 512" style="position: absolute; right:0">
                                        <path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"/>
                                    </svg>
                                </li>
                            </ul>
                            <input type="submit" value="Insert Category" onclick="CreateCategory(event)">
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>`
    $('#section-content').replaceWith(html)
}

function CreateCategory(event) {
    event.preventDefault()
    
    let category = {
        name: $('#name').val(),
        person: $('#person').val(),
        description: $('#description').val()
    }
    
    fetch('/admin/categories', {
        method: 'POST',
        headers: {'Content-Type' :'application/json'},
        body: JSON.stringify(category)
    })
    .then(response => response.json())
    .then(resJson => {
        if(!resJson.status) {
            alert("Create category failed! Error: " + resJson.error_msg)
            return
        }
        alert("Create category successed!")
        $('#name').val('')
        $('#description').val('')
    })
}

function GetEditCategory(id) {
    let url = '/admin/categories/' + id
    fetch(url, {
        method: 'GET',
    })
    .then(response => response.json())
    .then(data => {
        if(!data.status) {
            alert("Server Error: " + data.error_msg)
            return
        }

        category = data.data[0]
        let men =''
        let women=''
        let kids=''

        if(category.person === 'men') {
            men = 'selected'
        } else if (category.person === 'women') {
            women = 'selected'
        } else {
            kids = 'selected'
        }

        const html = `<section id="section-content">
        <div class="section-sub">
            <div class="section-title">
                <h4>Dashboard/ Update Categories</h4>
            </div>
            <div class="section-view">
                <div class="section-view-header">
                    <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20" viewBox="0 0 512 512">
                        <path d="M345 39.1L472.8 168.4c52.4 53 52.4 138.2 0 191.2L360.8 472.9c-9.3 9.4-24.5 9.5-33.9 .2s-9.5-24.5-.2-33.9L438.6 325.9c33.9-34.3 33.9-89.4 0-123.7L310.9 72.9c-9.3-9.4-9.2-24.6 .2-33.9s24.6-9.2 33.9 .2zM0 229.5V80C0 53.5 21.5 32 48 32H197.5c17 0 33.3 6.7 45.3 18.7l168 168c25 25 25 65.5 0 90.5L277.3 442.7c-25 25-65.5 25-90.5 0l-168-168C6.7 262.7 0 246.5 0 229.5zM144 144a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/>
                    </svg>
                    <h4>Update Category</h4>
                </div>
                <div class="section-view-content">
                    <div class="section-view-edit">
                        
                        <div class="section-form">
                            <form id="form-edit">
                                <ul>
                                    <li>
                                        <label for="id">Id</label>
                                        <input type="text" id="id" name="id" value="${category.id}" readonly>
                                    </li>
                                    <li style="position:relative">
                                        <label for="name">Category Name</label>
                                        <input type="text" id="name" name="name" value="${category.name}">
                                        <svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 512 512" style="position:absolute; right:0">
                                            <path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"/>
                                        </svg>
                                    </li>
                                    <li style="position:relative">
                                        <label for="person">Group</label>
                                        <select name="person" id="person" value=>
                                            <option value="men" ${men}>Men</option>
                                            <option value="women" ${women}}>Women</option>
                                            <option value="kids" ${kids}>Kids</option>
                                        </select>
                                    </li>
                                    <li style="position:relative">
                                        <label for="description">Description</label>
                                        <textarea type="text" id="description" rows="7" name="description">${category.description}</textarea>
                                        <svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 512 512" style="position: absolute; right:0">
                                            <path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.7 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"/>
                                        </svg>
                                    </li>
                                </ul>
                                <input type="submit" value="Update Category" id="btn-submit">
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>`
    $('#section-content').replaceWith(html)
    })

    
}