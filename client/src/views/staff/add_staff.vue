<template>
    <ion-page>
        <ion-content>
            <ion-grid>
                <ion-row>
                    <ion-col size="12">
                        <app-label>
                            <h4 v-if="!getStaffId">
                                {{ addStaffMessage }}
                            </h4>
                            <h4 v-if="getStaffId">
                                {{ updateStaffMessage }}
                            </h4>
                        </app-label>
                    </ion-col>
                </ion-row>
                <ion-row>
                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="staff-name"
                            type="text"
                            label="Staff Name"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :required="true"
                            v-model:input="staff.name"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <ion-select 
                            label="Role"
                            placeholder="Select Role"
                            fill="outline"
                            label-placement="floating"
                            v-model="staff.role"
                            @ionChange="handleRoleChange($event)"
                        >
                            <ion-select-option v-for="(role, index) in roles" :key="role + index" :value="role">
                                {{ role }}
                            </ion-select-option>
                        </ion-select>
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="staff-email"
                            type="email"
                            label="Email"
                            label-position="floating"
                            :counter="true"
                            :min-length="10"
                            :max-length="64"
                            :required="true"
                            v-model:input="staff.email"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4">
                        <app-input 
                            id="staff-password"
                            type="password"
                            label="Password"
                            label-position="floating"
                            :counter="true"
                            :min-length="8"
                            :max-length="16"
                            :required="true"
                            v-model:input="staff.password"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4" v-if="getStaffId">
                        <app-input 
                            id="staff-created-date"
                            type="text"
                            label="Created On"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :required="true"
                            v-model:input="staff.created_date"
                        />
                    </ion-col>

                    <ion-col size="8" size-md="4" size-lg="4" size-xl="2" size-xxl="4" v-if="getStaffId">
                        <app-input 
                            id="staff-modified-date"
                            type="text"
                            label="Modified On"
                            label-position="floating"
                            :counter="true"
                            :min-length="3"
                            :max-length="60"
                            :required="true"
                            v-model:input="staff.modified_date"
                        />
                    </ion-col>
                    
                </ion-row>
                <ion-row>
                    <ion-col size="12">
                        <app-switch
                            title="Status"
                            :checked="staff.status"
                            @change="onAppSwitchChange($event)"
                        />
                    </ion-col>
                </ion-row>
                <ion-row>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2">
                        <app-button 
                            id="clear-staff-btn"
                            title="Clear"
                            expand="block"
                            :action="clearStaff"
                        />
                    </ion-col>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2" v-if="!getStaffId">
                        <app-button 
                            id="add-staff-btn"
                            title="Save"
                            expand="block"
                            :action="createStaff"
                        />
                    </ion-col>
                    <ion-col size="6" size-md="4" size-lg="4" size-xl="2" size-xxl="2" v-if="getStaffId">
                        <app-button 
                            id="update-staff-btn"
                            title="Update"
                            expand="block"
                            :action="updateStaff"
                        />
                    </ion-col>
                </ion-row>
            </ion-grid>
        </ion-content>
    </ion-page>
</template>

<script>

import { IonCol, IonContent, IonGrid, IonPage, IonRow, IonSelect, IonSelectOption } from '@ionic/vue';
import AppLabel from '../../components/AppLabel.vue';
import AppInput from '../../components/AppInput.vue';
import AppButton from '../../components/AppButton.vue';
import AppSwitch from '../../components/AppSwitch.vue';
import { AddStaffMessage, UpdateStaffMessage, Roles } from '../../constants/constant.js';
import {
    mapActions,
    mapGetters
} from 'vuex';

export default {
    name: 'StaffAdd',
    components: {
        IonPage,
        IonContent,
        IonGrid,
        IonRow,
        IonCol,
        AppLabel,
        AppInput,
        AppButton,
        AppSwitch,
        IonSelect,
        IonSelectOption
    },
    data() {
        return {
            staffName: '',
            addStaffMessage: AddStaffMessage,
            updateStaffMessage: UpdateStaffMessage,
            staff: {
                status: true
            },
            roles: Roles
        }
    },
    computed: {
        ...mapGetters('staffStore', {
            getOpenedStaff: 'getOpenedStaff',
        }),
        ...mapGetters('authStore', {
            getLoggedInUser: 'getLoggedInUser'
        }),
        getStaffId() {
            const staffId = this.$route.params?.id;
            return staffId;
        }
    },
    methods: {
        ...mapActions('staffStore', {
            saveStaff: 'createStaff',
            getStaffById: 'getStaffById',
            updateStaffById: 'updateStaffById'
        }),
        onInput(event) {
            console.log(event);
            // api has be to be called 
        },
        async createStaff(event) {
            await this.saveStaff({
                data: {
                    ...this.staff,
                    bookshop_id: (this.getLoggedInUser?.user?._id || this.getLoggedInUser?.user?.id)
                },
                appOptions: {}
            });
            this.$router.push(`/add-staff/${this.getOpenedStaff?.InsertedID}`)
        },
        async updateStaff(event) {
            this.updateStaffById({
                data: {
                    ...this.staff,
                    bookshop_id: (this.getLoggedInUser?.user?._id || this.getLoggedInUser?.user?.id),
                    modified_date: new Date().toISOString()
                },
            });
        },
        clearStaff(event) {

        },
        onAppSwitchChange(event) {
            this.staff.status = !event.target.checked;
        },
        handleRoleChange(event) {
            
        }
    },
    async mounted() {
        if(this.$route.params.id) {
            await this.getStaffById({
                data: {
                    _id: this.$route.params.id
                }
            });
            Object.assign(this.staff, this.getOpenedStaff)
        }
    }
}

</script>