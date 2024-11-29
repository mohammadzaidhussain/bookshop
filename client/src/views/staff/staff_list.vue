<template>
    <ion-page>
        <ion-content>
            <ion-grid>
                <ion-row>
                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="search-staff"
                            type="text"
                            label="Search Staff"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            v-model:input="staffName"
                            @on-input="onInput($event)"
                        />
                    </ion-col>
                    <ion-col size="4" size-md="4" size-lg="4" size-xl="4" size-xxl="4">
                        <app-button 
                            id="add-staff-list-btn"
                            title="Add Staff"
                            :action="routeToAddStaff"
                        />
                    </ion-col>
                </ion-row>
                <ion-row>
                    <app-table
                        :columns="getColumnsConfig"
                        :data="getAllStaffs"
                        :pagination="false"
                        :total-records="100"
                        :limit="10"
                        :page-no="1"
                        :table-height="'60vh'"
                        @sort="onSort"
                        @load-page-no="onPageChange"
                        @on-limit-change="onLimitChange"
                    >
                        <template v-slot:action-buttons="{ rowData }">
                            <ion-icon :icon="eyeOutline" class="action-button bg-blue" @click="onActionIconClick('view', { data: rowData.row, index: rowData.rowIndex })"></ion-icon>
                            <ion-icon :icon="trashBinOutline" class="action-button bg-red" @click="onActionIconClick('delete', { data: rowData.row, index: rowData.rowIndex })"></ion-icon>
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
import storageService from '../../service/local_storage_service.js';


export default {
    name: 'StaffList',
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
            staffName: '',
            eyeOutline,
            trashBinOutline,
            columnsConfig: [
                {column_key: 'sno', column_name: 'S.No', sort: false, sortKey: 'sno', icon: '', showIcon: false, isVisible: true},
                {column_key: '_id', column_name: 'ID', sort: true, sortKey: 'sno', icon: '', showIcon: false, isVisible: true},
                {column_key: 'name', column_name: 'Name', sort: true, sortKey: 'name', icon: '', showIcon: true, isVisible: true},
                {column_key: 'email', column_name: 'Email', sort: true, sortKey: 'email', icon: '', showIcon: true, isVisible: true},
                {column_key: 'password', column_name: 'Password', sort: true, sortKey: 'password', icon: '', showIcon: false, isVisible: true},
                {column_key: 'role', column_name: 'Role', sort: true, sortKey: 'role', icon: '', showIcon: true, isVisible: true},
                {column_key: 'status', column_name: 'Status', sort: true, sortKey: 'status', icon: '', showIcon: true, isVisible: true},
                {column_key: 'created_date', column_name: 'Created On', sort: true, sortKey: 'created_date', icon: '', showIcon: true, isVisible: true},
                {column_key: 'action', column_name: 'Action', sort: false, sortKey: 'action', icon: '', showIcon: false, isVisible: true},
            ]
        }
    },
    computed: {
        ...mapGetters('staffStore', {
            getAllStaffs: 'getStaffs',
        }),
        ...mapGetters('authStore', {
            getLoggedInUser: 'getLoggedInUser'
        }),
        getColumnsConfig() {
            const cc = this.columnsConfig;
            return cc;
        }
    },
    methods: {
        ...mapActions('staffStore', {
            fetchAllStaffs: 'getAllStaffs',
            deleteStaffById: 'deleteStaffById'
        }),
        onInput(event) {
            console.log(event);
            // api has be to be called 
        },
        routeToAddStaff(event) {
            this.$router.push("/add-staff")
        },
        onSort(event) {
            console.log(event)
        },
        onPageChange(event) {
            console.log(event)
        },
        onLimitChange(event) {
            console.log(event)
        },
        onActionIconClick(operation, value) {
            switch (operation) {
                case "view":
                    this.$router.push(`/add-staff/${value.data._id || value.data.id}`)
                    break;
                case "delete":
                    this.deleteStaffById({
                        data: value.data
                    })
                    break;

                default:
                    break;
            }
        }
    },
    async mounted() {
        const userData = storageService.getPersistedData("logged_in_user");
        await this.fetchAllStaffs({
            data: {
                bookshop_id: userData?.user?._id
            }
        })
    }
}

</script>