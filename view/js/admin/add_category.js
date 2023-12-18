$(document).ready(function() {
    $('.success').hide()
    $('.fail').hide()

    $('.success button').click(function() {
        $('.success').hide()
    })
    $('.fail button').click(function() {
        $('.fail').hide()
    })
    $('#create_product').submit(function(event) {
        event.preventDefault()
        const fileInput = document.getElementById('imageurl')
        const file = fileInput.files[0]

        let title = $('#title').val()
        let summary = $('#summary').val()

        let jsonData = {
            id: "",
            name: title,
            description: summary,
            image: "",
        }
        let datajs = JSON.stringify(jsonData)

        const formData = new FormData();
        formData.append('file', file)
        formData.append('jsonData', datajs)
        fetch('/admin/categories', {
            method: 'POST',
            body: formData
        })
        .then(response => {
            $('.success').show()
            if(response.status === 200) {
                console.log("Them san pham thanh cong!")
                    $('.success').show()

            } else {
                console.log("Them san pham that bai!")
                    $('.fail').show()
            }
        })
    })
})