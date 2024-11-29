
const base_url = import.meta.env.VITE_API_URL;

export const auth_base_url = '/auth';
export const login_url = base_url + auth_base_url + '/login';
export const staff_login_url = base_url + auth_base_url + '/login/staff';

export const book_base_url = "/books";
export const create_book_url = base_url + book_base_url + "/create";
export const update_book_url = base_url + book_base_url + "/{id}/update";
export const delete_book_url = base_url + book_base_url + "/{id}/delete";
export const get_books_url = base_url + book_base_url + "/getAllBooks";
export const get_book_by_id_url = base_url + book_base_url + "/{id}"


export const staff_base_url = "/staffs";
export const create_staff_url = base_url + staff_base_url + "/create";
export const update_staff_url = base_url + staff_base_url + "/{id}/update";
export const delete_staff_url = base_url + staff_base_url + "/{id}/delete";
export const get_staffs_url = base_url + staff_base_url + "/getAllStaffs";
export const get_staff_by_id_url = base_url + staff_base_url + "/{id}";