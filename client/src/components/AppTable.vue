<template>
    <div class="table-container" :style="{'min-height': tableHeight}">
        <div class="app-table-wrapper">
            <div class="table-wrapper">
                <table class="app-table">
                    <colgroup ref="colgroup">
                        <col v-for="(column, index) in visibleColumns" :key="index" :style="{width: column.width || 'auto'}" />
                    </colgroup>
                    <thead>
                        <tr class="th-row">
                            <th v-for="(column, index) in visibleColumns" :key="index" @mousedown="startResize(index, $event)">
                                <div class="column-header">
                                    <app-label>
                                        <p>
                                            {{ column.column_name }}
                                        </p>
                                    </app-label>
                                    <ion-icon v-if="column.sort && column.sortKey" :icon="swapVerticalOutline" class="sort-icon" @click="handleSort(column)"></ion-icon>
                                </div>
                                <div class="resize-handle" @mousedown.stop="startResize(index, $event)"></div>
                            </th>
                        </tr>
                    </thead>
                    <tbody v-if="paginatedData?.length > 0">
                        <tr v-for="(row, rowIndex) in paginatedData" :key="rowIndex" :class="{'alt-row': rowIndex % 2 !== 0}">
                            <td v-for="(column, colIndex) in visibleColumns" :key="column.column_key+colIndex" :title="row[column.column_key]">
                                <span v-if="column.column_key == 'sno'" class="ellipsis-cell">
                                     {{ ((pageNo - 1) * limit) + (rowIndex + 1) }}
                                </span>
                                <span v-else-if="column.column_key == 'action'" class="ellipsis-cell">
                                    <slot name="action-buttons" :rowData="{ row, rowIndex }"></slot>
                                </span>
                                <span v-else class="ellipsis-cell">
                                    <app-label v-if="column.symbol">
                                        <p>
                                            {{ column.symbol }}
                                        </p>
                                    </app-label>
                                    <app-label>
                                        <p>
                                            {{ (row[column.column_key] === true) ? 'Active' : (row[column.column_key] === false) ? 'In Active' : row[column.column_key]?.toString() }}
                                        </p>
                                    </app-label>
                                </span>
                            </td>
                        </tr>
                    </tbody>
                    <div v-else>
                        No Records Found!
                    </div>
                </table>
            </div>
        </div>
        <div class="table-footer">
            <ion-grid>
                <ion-row>
                    <ion-col>
                        <div class="limit-selector">
                            <ion-item>
                                <label for="limit-select">show &nbsp;</label>
                                <select name="show" id="limit-select" @change="(event) => this.$emit('on-limit-change', {limit: event.target.value})">
                                    <option value="10">10</option>
                                    <option value="20">20</option>
                                    <option value="20">30</option>
                                    <option value="50">50</option>
                                </select>
                            </ion-item>
                        </div>
                    </ion-col>
                    <ion-col class="ion-justify-content-end">
                        <div v-if="pagination" :class="['pagination']">
                            <button v-for="page in totalPages" :key="page" class="page-btn" :class="{active: page == pageNo}" @click="loadPage(page)">
                                {{ page }}
                            </button>
                        </div>
                    </ion-col>
                </ion-row>
            </ion-grid>
        </div>
    </div>
</template>

<script>

import { IonCard, IonCardContent, IonCol, IonGrid, IonIcon, IonItem, IonRow } from '@ionic/vue';
import AppLabel from './AppLabel.vue';
import {
    swapVerticalOutline,
    eyeOutline,
    trashBinOutline
} from 'ionicons/icons';

export default {
    name: 'AppTable',
    components: {
        IonIcon,
        IonItem,
        IonGrid,
        IonRow,
        IonCol,
        IonCard,
        IonCardContent,
        AppLabel
    },
    props: {
        columns: {
            type: Array,
            required: true
        },
        data: {
            type: Array,
            required: true
        },
        pagination: {
            type: Boolean,
            default: false
        },
        totalRecords: {
            type: Number,
            default: 0
        },
        limit: {
            type: Number,
            default: 0
        },
        pageNo: {
            type: Number,
            default: 0
        },
        tableHeight: {
            type: String,
            default: '90vh'
        }
    },
    data() {
        return {
            sortKey: null,
            sortOrder: 'asc',
            resizingIndex: null,
            startX: 0,
            startWidth: 0,
            swapVerticalOutline,
            eyeOutline,
            trashBinOutline
        }
    },
    computed: {
        visibleColumns() {
            return this.columns.filter((column) => column.isVisible !== false)
        },
        paginatedData() {
            return this.data;
        },
        totalPages() {
            return Math.ceil(this.totalRecords / this.limit);
        },
    },
    methods: {
        handleSort(column) {
            if(!column.sort) return;
            this.sortKey = column.sortKey;
            this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc';
            this.$emit('sort', {
                sortKey: this.sortKey,
                sortOrder: this.sortOrder,
                pageNo: this.pageNo
            })
        },
        loadPage(page) {
            this.$emit('load-page-no', { pageNo: page });
        },
        handleResize(event) {
            if(this.resizingIndex !== null) {
                const newWidth = `${this.startWidth + event.clientX - this.startX}px`;
                this.visibleColumns[this.resizingIndex].width = newWidth;
            }
        },
        stopResize() {
            this.resizingIndex = null;
            document.addEventListener('mousemove', this.handleResize);
            document.addEventListener('mouseup', this.stopResize);
        },
        startResize(index, event) {
            this.resizingIndex = index;
            this.startX = event.clientX;
            this.startWidth = this.$refs.colgroup.children[index].offsetWidth;
            document.addEventListener('mousemove', this.handleResize);
            document.addEventListener('mouseup', this.stopResize);
        }
    }
}

</script>

<style scoped>

.table-container {
    overflow: auto;
    min-height: 70vh;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.app-table-wrapper {
    border-top-left-radius: 5px;
    border-top-right-radius: 5px;
}

.table-wrapper {
    overflow-x: auto;
}

.app-table {
    width: 100%;
    border: 1px solid rgb(230, 224, 224);
    border-collapse: collapse;
    table-layout: fixed;
}


.app-table th,
.app-table td {
    padding: 8px;
    text-align: left;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    min-width: 100px;
    max-width: 250px;
}

.app-table span.ellipsis-cell {
    display: flex;
}

.app-table span.ellipsis-cell:hover {
    white-space: normal;
    word-break: break-word;
    overflow: visible;
}

.app-table th {
    background-color: #5a5b5c;
    color: white;
    cursor: pointer;
    position: relative;
}

.column-header {
    display: flex;
    align-items: center;
}

.resize-handle {
    width: 5px;
    height: 100%;
    position: absolute;
    right: 0;
    top: 0;
    cursor: col-resize;
    border-right: 1px solid #ccc;
}

.sort-icon {
    margin-left: 5px;
    font-size: 14px;
}

.alt-row {
    background-color: #e5f3fa;
}

.pagination {
    margin-top: 10px;
    display: flex;
    justify-content: end;
    gap: 5px;
    padding: 5px;
}

.page-btn {
    padding: 5px 10px;
    border: 1px solid #ccc;
    cursor: pointer;
    background-color: #fff;
    border-radius: 50%;
    width: 35px;
    height: 35px;
    font-size: 0.7em;
}

.table-footer {
    display: flex;
}

.limit-selector {
    display:flex;
    justify-content: end;
}

#limit-select {
    width: 60px;
    height: 30px;
}

.active {
    background: #5a5b5c;
    color: white;
}

</style>