package types

const (
	//constants for replication session status
	RS_STATUS_UNKNOWN                               int = 0
	RS_STATUS_OTHER                                 int = 1
	RS_STATUS_OK                                    int = 2
	RS_STATUS_NON_RECOVERABLE_ERROR                 int = 7
	RS_STATUS_LOST_COMMUNICATION                    int = 13
	RS_STATUS_FAILED_OVER_WITH_SYNC                 int = 33792
	RS_STATUS_FAILED_OVER                           int = 33793
	RS_STATUS_MANUAL_SYNCING                        int = 33794
	RS_STATUS_PAUSED                                int = 33795
	RS_STATUS_IDLE                                  int = 33796
	RS_STATUS_AUTO_SYNC_CONFIGURED                  int = 33797
	RS_STATUS_DESTINATION_EXTEND_FAILED_NOT_SYNCING int = 33803
	RS_STATUS_DESTINATION_EXTEND_IN_PROGRESS        int = 33804
	RS_STATUS_ACTIVE                                int = 33805
	RS_STATUS_LOST_SYNC_COMMUNICATION               int = 33806
	RS_STATUS_DESTINATION_POOL_OUT_OF_SPACE         int = 33807
	RS_STATUS_SYNCING                               int = 33809
	RS_STATUS_FAILED_OVER_WITH_SYNC_MIXED           int = 34792
	RS_STATUS_FAILED_OVER_MIXED                     int = 34793
	RS_STATUS_MANUAL_SYNCING_MIXED                  int = 34794
	RS_STATUS_PAUSED_MIXED                          int = 34795
	RS_STATUS_IDLE_MIXED                            int = 34796
	RS_STATUS_AUTO_SYNC_CONFIGURED_MIXED            int = 34797
	RS_STATUS_HIBERNATED                            int = 34798
	RS_STATUS_ZERO_BANDWIDTH_CONFIGURED             int = 34799
	RS_STATUS_SOURCE_VDM_PERFORMANCE_DEGRADED       int = 34800
	RS_STATUS_SOURCE_RESOURCE_REPFORMANCE_DEGRADED  int = 34801
)
