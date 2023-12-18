$(document).ready(function() {
    $('.success').hide()
    $('.fail').hide()

    $('.success button').click(function() {
        $('.success').hide()
    })
    $('.fail button').click(function() {
        $('.fail').hide()
    })

    $("#Save").submit(function(event) {
        event.preventDefault()

        let id = $('#id').val()
        let pid = $('#p_id').val()
        let size = $('#size').val()
        let color = $('#color').val()
        let quantity = $('#quantity').val()

        let data = {
            id: parseInt(id),
            p_id: pid,
            size_product: parseInt(size),
            color: color,
            quantiry: parseInt(quantity)
        }
        let dataJson = JSON.stringify(data)
        
        const fileInput = document.getElementById('imageurl')
        const file = fileInput.files[0]

        const formData = new FormData();

        if(fileInput.files && fileInput.files.length > 0) {
            formData.append('file', file)
        }
        formData.append('jsonData', dataJson)

        let url = '/admin/products_version/update/' + id

        fetch(url, {
            method: 'PUT',
            body: formData
        })
        .then(response => {
            if (response.status === 200) {
                console.log("Thay doi san pham thanh cong")
                    $('.success').show()
            } else {
                console.log("Thay doi san pham that bai")
                    $('.fail').show()
            }
        })
    })
})