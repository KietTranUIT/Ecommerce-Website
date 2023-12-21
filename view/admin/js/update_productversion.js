$(document).ready(function() {
    $("#form-edit").submit(function(event) {
        event.preventDefault()

        let id = $('#id').val()
        let pid = $('#p_id').val()
        let size = $('#size').val()
        let quantity = $('#quantity').val()

        let data = {
            id: parseInt(id),
            p_id: parseInt(pid),
            size_product: parseInt(size),
            inventory: quantity
        }
        let dataJson = JSON.stringify(data)

        let url = '/admin/products_version/update/' + id

        fetch(url, {
            method: 'PUT',
            headers: {"Content-Type": "application/json"},
            body: dataJson
        })
        .then(response => {
            console.log(response)
            if (response.status === 200) {
                alert("Update product version successfully!")
            } else {
                alert("Update product version fail")
            }
        })
    })
})