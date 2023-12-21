$(document).ready(function() {
    $("#form-edit").submit(function(event) {
        event.preventDefault()

        let category_default = this.dataset.productId

        let id = $('#id').val()
        let name = $('#name').val()
        let description = $('#description').val()
        let category_id = $('#list').val()
        let price = $('#price').val()
        let image = $('#imageurl').val()

        if (category_id == '') {
            return
        } else {
            category_default = category_id
        }

        const fileInput = document.getElementById('image')
        const file = fileInput.files[0]

        let data = {
            id: parseInt(id),
            name: name,
            description: description,
            category_id: parseInt(category_default),
            price: parseInt(price),
            image: image,
        }
        let dataJson = JSON.stringify(data)

        let url = '/admin/products/update/' + id

        const formData = new FormData();
        formData.append('file', file)
        formData.append('jsonData', dataJson)
        fetch(url, {
            method: 'PUT',
            body: formData
        })
        .then(response => {
            if(response.status === 200) {
                alert('Update successfully')
            } else {
                alert('Update fail')
            }
        })
    })
})