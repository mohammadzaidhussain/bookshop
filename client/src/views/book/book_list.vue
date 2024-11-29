<template>
    <ion-page>
        <ion-content>
            <ion-grid>
                <ion-row>
                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="search-book"
                            type="text"
                            label="Search Book"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :debounce="debounce"
                            v-model:input="bookName"
                            @on-input="onInput($event)"
                        />
                    </ion-col>
                    <ion-col size="4" size-md="4" size-lg="4" size-xl="4" size-xxl="4">
                        <app-button 
                            id="add-book-list-btn"
                            title="Add Book"
                            :action="routeToAddBook"
                        />
                    </ion-col>
                </ion-row>
                <ion-row>
                    <app-table
                        :columns="getColumnsConfig"
                        :data="getAllBooks"
                        :pagination="true"
                        :total-records="getTotalBookRecords"
                        :limit="limit"
                        :page-no="pageNo"
                        :table-height="'60vh'"
                        @sort="onSort"
                        @load-page-no="onPageChange"
                        @on-limit-change="onLimitChange"
                    >
                        <template v-slot:action-buttons="{ rowData }">
                            <ion-icon :icon="eyeOutline" class="action-button bg-blue" @click="onActionIconClick('view', { data: rowData.row, index: rowData.rowIndex })"></ion-icon>
                            <ion-icon v-if="getLoggedInUser?.role != 'staff'" :icon="trashBinOutline" class="action-button bg-red" @click="onActionIconClick('delete', { data: rowData.row, index: rowData.rowIndex })"></ion-icon>
                        </template>
                    </app-table>
                </ion-row>
            </ion-grid>
        </ion-content>
    </ion-page>
</template>

<script>

import { IonCol, IonContent, IonGrid, IonIcon, IonPage, IonRow } from '@ionic/vue';
import AppLabel from '../../components/AppLabel.vue';
import AppInput from '../../components/AppInput.vue';
import AppButton from '../../components/AppButton.vue';
import AppTable from '../../components/AppTable.vue';
import { eyeOutline, trashBinOutline } from 'ionicons/icons';
import {
    mapGetters,
    mapActions
} from "vuex";


export default {
    name: 'BookList',
    components: {
        IonPage,
        IonContent,
        IonGrid,
        IonRow,
        IonCol,
        AppLabel,
        AppInput,
        AppButton,
        AppTable,
        IonIcon
    },
    data() {
        return {
            bookName: '',
            eyeOutline,
            trashBinOutline,
            columnsConfig: [ 
                {column_key: 'sno', column_name: 'S.No', sort: false, sortKey: 'sno', icon: '', showIcon: false, isVisible: true},
                {column_key: '_id', column_name: 'ID', sort: true, sortKey: 'sno', icon: '', showIcon: false, isVisible: true},
                {column_key: 'book_name', column_name: 'Name', sort: true, sortKey: 'book_name', icon: '', showIcon: true, isVisible: true},
                {column_key: 'book_code', column_name: 'Code', sort: true, sortKey: 'book_code', icon: '', showIcon: true, isVisible: true},
                {column_key: 'book_price', column_name: 'Price', sort: true, sortKey: 'book_price', icon: '', showIcon: true, isVisible: true},
                {column_key: 'author', column_name: 'Author', sort: true, sortKey: 'author', icon: '', showIcon: true, isVisible: true},
                {column_key: 'published_year', column_name: 'Published Year', sort: true, sortKey: 'published_year', icon: '', showIcon: true, isVisible: true},
                {column_key: 'status', column_name: 'Status', sort: true, sortKey: 'status', icon: '', showIcon: true, isVisible: true},
                {column_key: 'created_date', column_name: 'Created On', sort: true, sortKey: 'created_date', icon: '', showIcon: true, isVisible: true},
                {column_key: 'action', column_name: 'Action', sort: false, sortKey: 'action', icon: '', showIcon: false, isVisible: true},
            ],
            pageNo: 1,
            limit: 10,
            debounce: 1000
        }
    },
    watch: {
        bookName(newValue) {
            this.fetchAllBooks({
                data: {
                    page_no: parseInt(this.pageNo),
                    limit: parseInt(this.limit),
                    book_search: newValue
                }
            });
        }
    },
    computed: {
        ...mapGetters('bookStore', {
            getAllBooks: 'getBooks',
            getTotalBookRecords: 'getTotalBookRecords'
        }),
        ...mapGetters('authStore', {
            getLoggedInUser: 'getLoggedInUser',
        }),
        getColumnsConfig() {
            const cc = this.columnsConfig;
            const priceColIndex = this.columnsConfig.findIndex((c) => c.column_key == "book_price");
            if(priceColIndex > -1) {
                cc[priceColIndex].symbol = "$";
            }
            return cc;
        }
    },
    methods: {
        ...mapActions('bookStore', {
            fetchAllBooks: 'getAllBooks',
            deleteBookById: 'deleteBookById'
        }),
        onInput(event) {
            console.log(event);
            // api has be to be called 
        },
        routeToAddBook(event) {
            this.$router.push("/add-book")
        },
        onSort(event) {
            console.log(event)
        },
        async onPageChange(event) {
            this.pageNo = event.pageNo;
            this.$router.push(`/books?page_no=${this.pageNo}&limit=${this.limit}`)
            await this.fetchAllBooks({
                data: {
                    page_no: parseInt(this.pageNo),
                    limit: parseInt(this.limit)
                }
            });
        },
        async onLimitChange(event) {
            this.limit = event.limit;
            this.$router.push(`/books?page_no=${this.pageNo}&limit=${this.limit}`)
            await this.fetchAllBooks({
                data: {
                    page_no: parseInt(this.pageNo),
                    limit: parseInt(this.limit)
                }
            });
        },
        onActionIconClick(operation, value) {
            switch (operation) {
                case "view":
                    this.$router.push(`/add-book/${value.data._id || value.data.id}`)
                    break;
                case "delete":
                    this.deleteBookById({
                        data: value.data
                    })
                    break;

                default:
                    break;
            }
        }
    },
    async mounted() {
        if(this.$route.query?.page_no) {
            this.pageNo = this.$route.query?.page_no;
        }
        if(this.$route.query?.limit) {
            this.limit = this.$route.query?.limit;
        }
        await this.fetchAllBooks({
            data: {
                page_no: parseInt(this.pageNo),
                limit: parseInt(this.limit)
            }
        })
    }
}

</script>