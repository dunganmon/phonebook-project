
function getAuthHeader() {
    const auth = sessionStorage.getItem('auth');
    if (!auth) {
        window.location.href = 'login.html';
        return null;
    }
    return { 'Authorization': 'Basic ' + auth };
}


function checkAuth() {
    const auth = sessionStorage.getItem('auth');
    if (!auth) {
        window.location.href = 'login.html';
        return null;
    }
    return sessionStorage.getItem('username');
}


function loadUserName() {
    const username = checkAuth();
    if (username) {
        const userNameElement = document.getElementById('userName');
        if (userNameElement) userNameElement.textContent = username;
    }
}


function logout() {
    sessionStorage.clear();
    window.location.href = 'login.html';
}


async function loadContacts() {
    const headers = getAuthHeader();
    if (!headers) return;

    try {
        const res = await fetch('http://localhost:8080/api/contacts', {
            method: 'GET',
            headers: headers
        });
        if (!res.ok) throw new Error('Lỗi khi tải danh bạ');
        const data = await res.json();
        const contacts = data.data.contacts;

        const contactsTableBody = document.getElementById('contactsTableBody');
        contactsTableBody.innerHTML = '';

        contacts.forEach(contact => {
            const newRow = document.createElement('tr');
            newRow.innerHTML = `
                <td>${contact.full_name}</td>
                <td>${contact.phone}</td>
                <td>${contact.email || ''}</td>
                <td>${contact.address || ''}</td>
                <td class="action-buttons">
                    <button class="edit-btn" data-id="${contact.id}">Sửa</button>
                    <button class="delete-btn" data-id="${contact.id}">Xóa</button>
                </td>
            `;
            contactsTableBody.appendChild(newRow);
        });
    } catch (error) {
        alert(error.message);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    loadUserName();
    loadContacts();

    const addContactForm = document.getElementById('addContactForm');
    const contactsTableBody = document.getElementById('contactsTableBody');

 
    addContactForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const headers = getAuthHeader();
        if (!headers) return;

        const fullName = document.getElementById('name').value.trim();
        const phone = document.getElementById('tel').value.trim();
        const email = document.getElementById('email').value.trim();
        const address = document.getElementById('address').value.trim();

        if (!fullName || !phone) {
            alert('Vui lòng điền đầy đủ Họ tên và Số điện thoại!');
            return;
        }

        try {
            const res = await fetch('http://localhost:8080/api/contacts', {
                method: 'POST',
                headers: {
                    ...headers,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ full_name: fullName, phone, email, address })
            });
            if (!res.ok) throw new Error('Thêm liên hệ thất bại');
            await loadContacts();
            addContactForm.reset();
            alert('Thêm liên hệ thành công');
        } catch (error) {
            alert(error.message);
        }
    });

   
    contactsTableBody.addEventListener('click', async (event) => {
        const headers = getAuthHeader();
        if (!headers) return;

        if (event.target.classList.contains('delete-btn')) {
            const id = event.target.dataset.id;
            if (!confirm('Bạn có chắc muốn xóa liên hệ này?')) return;

            try {
                const res = await fetch(`http://localhost:8080/api/contacts/${id}`, {
                    method: 'DELETE',
                    headers: headers
                });
                if (!res.ok) throw new Error('Xóa liên hệ thất bại');
                await loadContacts();
                alert('Xóa liên hệ thành công');
            } catch (error) {
                alert(error.message);
            }
        }

        if (event.target.classList.contains('edit-btn')) {
            const row = event.target.closest('tr');
            const id = event.target.dataset.id;

            if (event.target.textContent === 'Sửa') {
                for (let i = 0; i < 4; i++) {
                    const cell = row.cells[i];
                    const text = cell.textContent;
                    cell.innerHTML = `<input type="text" value="${text}">`;
                }
                event.target.textContent = 'Lưu';
            } else {
                const inputs = row.querySelectorAll('input');
                const updatedContact = {
                    full_name: inputs[0].value.trim(),
                    phone: inputs[1].value.trim(),
                    email: inputs[2].value.trim(),
                    address: inputs[3].value.trim()
                };

                try {
                    const res = await fetch(`http://localhost:8080/api/contacts/${id}`, {
                        method: 'PUT',
                        headers: {
                            ...headers,
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(updatedContact)
                    });
                    if (!res.ok) throw new Error('Cập nhật liên hệ thất bại');
                    await loadContacts();
                    alert('Cập nhật liên hệ thành công');
                } catch (error) {
                    alert(error.message);
                }
                event.target.textContent = 'Sửa';
            }
        }
    });
});
