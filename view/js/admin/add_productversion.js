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

        let p_id = $('#p_id').val()
        let size = $('#size').val()
        let color = $('#color').val()
        let quantity = $('#inventory').val()

        let jsonData = {
            id: "",
            p_id: parseInt(p_id),
            size_product: parseInt(size),
            color: color,
            image: "",
            inventory: quantity 
        }
        let datajs = JSON.stringify(jsonData)
        console.log(datajs)

        const formData = new FormData();
        formData.append('file', file)
        formData.append('jsonData', datajs)

        let url = '/admin/products/' + p_id + '/products_version' 
        fetch(url, {
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