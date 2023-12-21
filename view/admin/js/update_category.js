$(document).ready(function() {
    $("#form-edit").submit(function(event) {
        event.preventDefault()

        let id = $('#id').val()
        let name = $('#name').val()
        let description = $('#description').val()
        let imageurl = $('#imageurl').val()

        let data = {
            id: id,
            name: name,
            description: description,
            image: imageurl
        }
        let dataJson = JSON.stringify(data)
        
        const fileInput = document.getElementById('image')
        const file = fileInput.files[0]

        const formData = new FormData();

        if(fileInput.files && fileInput.files.length > 0) {
            formData.append('file', file)
        }
        formData.append('jsonData', dataJson)

        let url = '/admin/categories/update/' + id

        fetch(url, {
            method: 'PUT',
            body: formData
        })
        .then(response => {
            if (response.status === 200) {
                alert("Update Category successfully")
            } else {
                alert("Update Category failed")
            }
        })
    })
})