
// Get list category
export function GetCategory() {
    fetch('/admin/categories', {
        method: 'GET'
    })
    .then(response => response.json())
    .then(data => {
        if(data.status === false) {
            alert("Server error: " + data.error_msg)
            return
        }

        let categories = data.data[0]
        let categoriesHTML = ''
        categories.forEach(category => {
            const content = `
            <tr>
                <td>${category.id}</td>
                <td>${category.name}</td>
                <td>${category.person}</td>
                <td>
                    <p style="width:800px;overflow:hidden;text-overflow:ellipsis">${category.description}</p>
                </td>
                <td>
                    <button type="button" data-category-id="${category.id}" class="btn btn-primary btn-edit" onclick="GetEditCategory(${category.id})">Edit</button>
                    <button type="button" data-category-id="${category.id}" class="btn btn-outline-secondary btn-detail" onclick="">Details</button>
                    <button type="button" data-category-id="${category.id}" class="btn btn-danger btn-delete">Delete</button>
                </td>
            </tr>`
            categoriesHTML += content
        })

        const html = `<section id="section-content">
            <div class="section-sub">
                <div class="section-title">
                    <h4>Dashboard/ View Categories</h4>
                    <button type="button" class="btn btn-primary" onclick="GetNewCategory()">Create New</button>
                </div>
                <div class="section-view">
                    <div class="section-view-header">
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20" viewBox="0 0 512 512">
                            <path d="M345 39.1L472.8 168.4c52.4 53 52.4 138.2 0 191.2L360.8 472.9c-9.3 9.4-24.5 9.5-33.9 .2s-9.5-24.5-.2-33.9L438.6 325.9c33.9-34.3 33.9-89.4 0-123.7L310.9 72.9c-9.3-9.4-9.2-24.6 .2-33.9s24.6-9.2 33.9 .2zM0 229.5V80C0 53.5 21.5 32 48 32H197.5c17 0 33.3 6.7 45.3 18.7l168 168c25 25 25 65.5 0 90.5L277.3 442.7c-25 25-65.5 25-90.5 0l-168-168C6.7 262.7 0 246.5 0 229.5zM144 144a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/>
                        </svg>
                        <h4>View Categories</h4>
                    </div>
                    <div class="section-view-content">
                        <div class="section-view-table">
                            <table>
                                <colgroup>
                                    <col>
                                    <col>
                                    <col>
                                    <col>
                                    <col>
                                </colgroup>
                                <thead>
                                    <th>Id</th>
                                    <th>Name</th>
                                    <th>For</th>
                                    <th>Description</th>
                                    <th>Edit</th>
                                </thead>
                                <tbody>
                                    ${categoriesHTML}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </section>`
        $('#section-content').replaceWith(html)

        history.pushState(null, '/admin/home', '/admin/categories')
    })
}

export function CreateCategory() {
    $('#form-insert').submit(function(event) {
        event.preventDefault()
        const fileInput = document.getElementById('image')
        const file = fileInput.files[0]
    
        let title = $('#name').val()
        let summary = $('#description').val()
    
        let jsonData = {
            id: "",
            name: title,
            description: summary,
            image: "",
        }
        let datajs = JSON.stringify(jsonData)
        console.log(datajs)
    
        const formData = new FormData();
        formData.append('file', file)
        formData.append('jsonData', datajs)
        fetch('/admin/categories', {
            method: 'POST',
            body: formData
        })
        .then(response => {
            if(response.status === 200) {
                alert('Add category successfull')
            } else {
                alert('Add category fail')
            }
        })
    })
}