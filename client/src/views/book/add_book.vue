<template>
    <ion-page>
        <ion-content>
            <ion-grid>
                <ion-row>
                    <ion-col size="12">
                        <app-label>
                            <h4 v-if="!getBookId">
                                {{ addBookMessage }}
                            </h4>
                            <h4 v-if="getBookId">
                                {{ updateBookMessage }}
                            </h4>
                        </app-label>
                    </ion-col>
                </ion-row>
                <ion-row>
                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-name"
                            type="text"
                            label="Book Name"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :required="true"
                            v-model:input="book.book_name"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-code"
                            type="text"
                            label="Book Code"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="6"
                            :required="true"
                            v-model:input="book.book_code"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-price"
                            type="number"
                            label="Price"
                            label-position="floating"
                            :counter="true"
                            :min-length="1"
                            :max-length="5"
                            :required="true"
                            v-model:input="book.book_price"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-page"
                            type="number"
                            label="Page Count"
                            label-position="floating"
                            :counter="true"
                            :min-length="2"
                            :max-length="4"
                            :required="true"
                            v-model:input="book.page_count"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-author"
                            type="text"
                            label="Author"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :required="true"
                            v-model:input="book.author"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="book-author"
                            type="number"
                            label="Published Year"
                            label-position="floating"
                            :counter="true"
                            :min-length="4"
                            :max-length="4"
                            :required="true"
                            v-model:input="book.published_year"
                        />
                    </ion-col>
                    
                </ion-row>
                <ion-row>
                    <ion-col size="12">
                        <app-input 
                            id="book-description"
                            type="text"
                            label="Book Description"
                            label-position="floating"
                            :counter="true"
                            :min-length="0"
                            :max-length="1024"
                            :required="true"
                            v-model:input="book.description"
                        />
                    </ion-col>
                </ion-row>
                <ion-row>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2">
                        <app-button 
                            id="clear-book-btn"
                            title="Clear"
                            expand="block"
                            :action="clearBook"
                        />
                    </ion-col>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2" v-if="!getBookId">
                        <app-button 
                            id="add-book-btn"
                            title="Save"
                            expand="block"
                            :action="createBook"
                        />
                    </ion-col>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2" v-if="getBookId">
                        <app-button 
                            id="update-book-btn"
                            title="Update"
                            expand="block"
                            :action="updateBook"
                        />
                    </ion-col>
                </ion-row>
            </ion-grid>
        </ion-content>
    </ion-page>
</template>

<script>

import { IonCol, IonContent, IonGrid, IonPage, IonRow } from '@ionic/vue';
import AppLabel from '../../components/AppLabel.vue';
import AppInput from '../../components/AppInput.vue';
import AppButton from '../../components/AppButton.vue';
import { AddBookMessage, UpdateBookMessage } from '../../constants/constant.js';
import {
    mapActions,
    mapGetters
} from 'vuex';

export default {
    name: 'BookAdd',
    components: {
        IonPage,
        IonContent,
        IonGrid,
        IonRow,
        IonCol,
        AppLabel,
        AppInput,
        AppButton
    },
    data() {
        return {
            bookName: '',
            addBookMessage: AddBookMessage,
            updateBookMessage: UpdateBookMessage,
            book: {}
        }
    },
    computed: {
        ...mapGetters('bookStore', {
            getOpenedBook: 'getOpenedBook',
        }),
        getBookId() {
            const bookId = this.$route.params?.id;
            return bookId;
        }
    },
    methods: {
        ...mapActions('bookStore', {
            saveBook: 'createBook',
            getBookById: 'getBookById',
            updateBookById: 'updateBookById'
        }),
        onInput(event) {
            console.log(event);
            // api has be to be called 
        },
        async createBook(event) {
            await this.saveBook({
                data: {
                    ...this.book,
                    book_price: parseFloat(this.book.book_price || 0),
                    page_count: parseFloat(this.book.page_count || 0),
                    published_year: parseFloat(this.book.published_year || 0),
                },
                appOptions: {}
            });
            this.$router.push(`/add-book/${this.getOpenedBook?.InsertedID}`)
        },
        async updateBook(event) {
            this.updateBookById({
                data: {
                    ...this.book,
                    book_price: parseFloat(this.book.book_price || 0),
                    page_count: parseFloat(this.book.page_count || 0),
                    published_year: parseFloat(this.book.published_year || 0),
                    modified_date: new Date().toISOString()
                },
            });
        },
        clearBook(event) {

        }
    },
    async mounted() {
        if(this.$route.params.id) {
            await this.getBookById({
                data: {
                    _id: this.$route.params.id
                }
            });
            Object.assign(this.book, this.getOpenedBook)
        }
    }
}

</script>