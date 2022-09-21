package postgres13

import (
	"github.com/wader/fq/format/postgres/common"
	"github.com/wader/fq/format/postgres/flavours/postgres14/common14"
	"github.com/wader/fq/pkg/decode"
)

// type = struct PageHeaderData {
/*    0      |     8 */ // PageXLogRecPtr pd_lsn;
/*    8      |     2 */ // uint16 pd_checksum;
/*   10      |     2 */ // uint16 pd_flags;
/*   12      |     2 */ // LocationIndex pd_lower;
/*   14      |     2 */ // LocationIndex pd_upper;
/*   16      |     2 */ // LocationIndex pd_special;
/*   18      |     2 */ // uint16 pd_pagesize_version;
/*   20      |     4 */ // TransactionId pd_prune_xid;
/*   24      |     0 */ // ItemIdData pd_linp[];
//
/* total size (bytes):   24 */

// type = struct HeapTupleHeaderData {
/*    0      |    12 */ // union {
/*                12 */ //    HeapTupleFields t_heap;
/*                12 */ //    DatumTupleFields t_datum;
//					       } t_choice;
/* total size (bytes):   12 */
//
/*   12      |     6 */ // ItemPointerData t_ctid
/*   18      |     2 */ // uint16 t_infomask2
/*   20      |     2 */ // uint16 t_infomask
/*   22      |     1 */ // uint8 t_hoff
/*   23      |     0 */ // bits8 t_bits[]
/* XXX  1-byte padding  */
//
/* total size (bytes):   24 */

// type = struct HeapTupleFields {
/*    0      |     4 */ // TransactionId t_xmin;
/*    4      |     4 */ // TransactionId t_xmax;
/*    8      |     4 */ // union {
/*                 4 */ //    CommandId t_cid;
/*                 4 */ //    TransactionId t_xvac;
//                         } t_field3;h
/* total size (bytes):    4 */
//
/* total size (bytes):   12 */

func DecodeHeap(d *decode.D) any {
	heap := &common14.HeapD{
		PageSize:               common.HeapPageSize,
		DecodePageHeaderDataFn: common14.DecodePageHeader,
	}
	return common14.DecodeHeap(heap, d)
}
