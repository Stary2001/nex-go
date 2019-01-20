// This file is autogenerated.
// I apologise in advance.
package nex

type PID uint32
type Result uint32
type Buffer []byte
type QBuffer []byte
type DateTime uint64
type StationURL string
type AccountData struct {
    Base Structure
    Pid PID
    StrName string
    UiGroups uint32
    StrEmail string
    DtCreationDate DateTime
    DtEffectiveDate DateTime
    StrNotEffectiveMsg string
    DtExpiryDate DateTime
    StrExpiredMsg string
}

type AutoMatchmakeParam struct {
    Base Structure
    SourceMatchmakeSession MatchmakeSession
    AdditionalParticipants []PID
    GidForParticipationCheck uint32
    AutoMatchmakeOption uint32
    JoinMessage string
    ParticipationCount uint16
    LstSearchCriteria []MatchmakeSessionSearchCriteria
    TargetGids []uint32
    BlockListParam MatchmakeBlockListParam
}

type BasicAccountInfo struct {
    Base Structure
    PidOwner PID
    StrName string
}

type BlacklistedPrincipal struct {
    PrincipalBasicInfo PrincipalBasicInfo
    GameKey GameKey
    BlacklistedSince DateTime
}

type BufferQueueParam struct {
    Base Structure
    DataId uint64
    Slot uint32
}

type Comment struct {
    Unknown uint8
    StatusMessage string
    LastChangedOn DateTime
}

type ConnectionData struct {
    Base Structure
    StationUrl StationURL
    ConnectionID uint32
}

type CreateMatchmakeSessionParam struct {
    Base Structure
    SourceMatchmakeSession MatchmakeSession
    AdditionalParticipants []PID
    GidForParticipationCheck uint32
    CreateMatchmakeSessionOption uint32
    JoinMessage string
    ParticipationCount uint16
}

type Data struct {
    type_name string
    len_plus_four uint32
    data Buffer
}

type DataStoreChangeMetaCompareParam struct {
    Base Structure
    ComparisonFlag uint32
    Name string
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary QBuffer
    Tags []string
    ReferredCnt uint32
    DataType uint16
    Status uint8
}

type DataStoreChangeMetaParam struct {
    Base Structure
    DataId uint64
    ModifiesFlag uint32
    Name string
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary QBuffer
    Tags []string
    UpdatePassword uint64
    ReferredCnt uint32
    DataType uint16
    Status uint8
    CompareParam DataStoreChangeMetaCompareParam
    PersistenceTarget DataStorePersistenceTarget
}

type DataStoreChangeMetaParamV1 struct {
    Base Structure
    DataId uint64
    ModifiesFlag uint32
    Name string
    Permission DataStorePermission
    DelPermission DataStorePermission
    Period uint16
    MetaBinary QBuffer
    Tags []string
    UpdatePassword uint64
}

type DataStoreCompletePostParam struct {
    Base Structure
    DataId uint64
    IsSuccess bool
}

type DataStoreCompletePostParamV1 struct {
    Base Structure
    DataId uint32
    IsSuccess bool
}

type DataStoreCompleteUpdateParam struct {
    Base Structure
    DataId uint64
    Version uint32
    IsSuccess bool
}

type DataStoreDeleteParam struct {
    Base Structure
    DataId uint64
    UpdatePassword uint64
}

type DataStoreFetchMyInfosAchievementResult struct {
    Base Structure
    DataId uint64
    DataType uint16
    MetaBinary QBuffer
    CreatedTime DateTime
    Ratings map[int8]DataStoreRatingInfo
    Buffers map[int8][]QBuffer
}

type DataStoreFetchMyInfosBalloonResult struct {
    Base Structure
    DataId uint64
    DataType uint16
    MetaBinary QBuffer
    CreatedTime DateTime
    UpdatedTime DateTime
    IsCleared bool
    Ratings map[int8]DataStoreRatingInfo
    Buffers map[int8][]QBuffer
}

type DataStoreFetchMyInfosParam struct {
    Base Structure
    BalloonDataTypes []uint16
    AdditionalOperation uint16
}

type DataStoreFetchMyInfosResult struct {
    Base Structure
    Balloons []DataStoreFetchMyInfosBalloonResult
    Achievement DataStoreFetchMyInfosAchievementResult
}

type DataStoreGetMetaParam struct {
    Base Structure
    DataId uint64
    PersistenceTarget DataStorePersistenceTarget
    ResultOption uint8
    AccessPassword uint64
}

type DataStoreGetNewArrivedNotificationsParam struct {
    Base Structure
    LastNotificationId uint64
    Limit uint16
}

type DataStoreGetNotificationUrlParam struct {
    Base Structure
    PreviousUrl string
}

type DataStoreGetSpecificMetaParam struct {
    Base Structure
    DataIds []uint64
}

type DataStoreGetSpecificMetaParamV1 struct {
    Base Structure
    DataIds []uint32
}

type DataStoreKeyValue struct {
    Base Structure
    Key string
    Value string
}

type DataStoreMetaInfo struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    Name string
    DataType uint16
    MetaBinary QBuffer
    Permission DataStorePermission
    DelPermission DataStorePermission
    CreatedTime DateTime
    UpdatedTime DateTime
    Period uint16
    Status uint8
    ReferredCnt uint32
    ReferDataId uint32
    Flag uint32
    ReferredTime DateTime
    ExpireTime DateTime
    Tags []string
    Ratings []DataStoreRatingInfoWithSlot
}

type DataStoreNotification struct {
    Base Structure
    NotificationId uint64
    DataId uint64
}

type DataStoreNotificationV1 struct {
    Base Structure
    NotificationId uint64
    DataId uint32
}

type DataStorePasswordInfo struct {
    Base Structure
    DataId uint64
    AccessPassword uint64
    UpdatePassword uint64
}

type DataStorePermission struct {
    Base Structure
    Permission uint8
    RecipientIds []uint64
}

type DataStorePersistenceInfo struct {
    Base Structure
    OwnerId uint64
    PersistenceSlotId uint16
    DataId uint64
}

type DataStorePersistenceInitParam struct {
    Base Structure
    PersistenceSlotId uint16
    DeleteLastObject bool
}

type DataStorePersistenceTarget struct {
    Base Structure
    OwnerId uint64
    PersistenceSlotId uint16
}

type DataStorePrepareGetParam struct {
    Base Structure
    DataId uint64
    LockId uint32
    PersistenceTarget DataStorePersistenceTarget
    AccessPassword uint64
    ExtraData []string
}

type DataStorePrepareGetParamV1 struct {
    Base Structure
    DataId uint32
    LockId uint32
}

type DataStorePreparePostParam struct {
    Base Structure
    Size uint32
    Name string
    DataType uint16
    MetaBinary QBuffer
    Permission DataStorePermission
    DelPermission DataStorePermission
    Flag uint32
    Period uint16
    ReferDataId uint32
    Tags []string
    RatingInitParams []DataStoreRatingInitParamWithSlot
    PersistenceInitParam DataStorePersistenceInitParam
    ExtraData []string
}

type DataStorePreparePostParamV1 struct {
    Base Structure
    Size uint32
    Name string
    DataType uint16
    MetaBinary QBuffer
    Permission DataStorePermission
    DelPermission DataStorePermission
    Flag uint32
    Period uint16
    ReferDataId uint32
    Tags []string
    RatingInitParams []DataStoreRatingInitParamWithSlot
}

type DataStorePrepareUpdateParam struct {
    Base Structure
    DataId uint64
    Size uint32
    UpdatePassword uint64
    ExtraData []string
}

type DataStoreRateObjectParam struct {
    Base Structure
    RatingValue int32
    AccessPassword uint64
}

type DataStoreRatingInfo struct {
    Base Structure
    TotalValue int64
    Count uint32
    InitialValue int64
}

type DataStoreRatingInfoWithSlot struct {
    Base Structure
    Slot int8
    Rating DataStoreRatingInfo
}

type DataStoreRatingInitParam struct {
    Base Structure
    Flag uint8
    InternalFlag uint8
    LockType uint8
    InitialValue int64
    RangeMin int32
    RangeMax int32
    PeriodHour int8
    PeriodDuration int16
}

type DataStoreRatingInitParamWithSlot struct {
    Base Structure
    Slot int8
    Param DataStoreRatingInitParam
}

type DataStoreRatingLog struct {
    Base Structure
    IsRated bool
    Pid uint64
    RatingValue int32
    LockExpirationTime DateTime
}

type DataStoreRatingTarget struct {
    Base Structure
    DataId uint64
    Slot int8
}

type DataStoreReqGetAdditionalMeta struct {
    Base Structure
    OwnerId uint64
    DataType uint16
    Version uint16
    MetaBinary QBuffer
}

type DataStoreReqGetInfo struct {
    Base Structure
    Url string
    RequestHeaders []DataStoreKeyValue
    Size uint32
    RootCaCert Buffer
    DataId uint64
}

type DataStoreReqGetInfoV1 struct {
    Base Structure
    Url string
    RequestHeaders []DataStoreKeyValue
    Size uint32
    RootCaCert Buffer
}

type DataStoreReqGetNotificationUrlInfo struct {
    Base Structure
    Url string
    Key string
    Query string
    RootCaCert Buffer
}

type DataStoreReqPostInfo struct {
    Base Structure
    DataId uint64
    Url string
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert Buffer
}

type DataStoreReqPostInfoV1 struct {
    Base Structure
    DataId uint32
    Url string
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert Buffer
}

type DataStoreReqUpdateInfo struct {
    Base Structure
    Version uint32
    Url string
    RequestHeaders []DataStoreKeyValue
    FormFields []DataStoreKeyValue
    RootCaCert Buffer
}

type DataStoreSearchBalloonParam struct {
    Base Structure
    DataType uint16
    UserRank uint8
    ResultSetCount uint8
}

type DataStoreSearchBalloonResult struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    Name string
    DataType uint16
    MetaBinary QBuffer
    CreatedTime DateTime
    UpdatedTime DateTime
    OwnerDataId uint64
    OwnerName string
    IsFriendBalloon bool
    Ratings map[int8]DataStoreRatingInfo
    OwnerRatings map[int8]DataStoreRatingInfo
}

type DataStoreSearchBalloonResultSet struct {
    Base Structure
    Balloons []DataStoreSearchBalloonResult
}

type DataStoreSearchParam struct {
    Base Structure
    SearchTarget uint8
    OwnerIds []uint64
    OwnerType uint8
    DestinationIds []uint64
    DataType uint16
    CreatedAfter DateTime
    CreatedBefore DateTime
    UpdatedAfter DateTime
    UpdatedBefore DateTime
    ReferDataId uint32
    Tags []string
    ResultOrderColumn uint8
    ResultOrder uint8
    ResultRange ResultRange
    ResultOption uint8
    MinimalRatingFrequency uint32
    UseCache bool
    TotalCountEnabled bool
    DataTypes []uint16
}

type DataStoreSearchResult struct {
    Base Structure
    TotalCount uint32
    Result []DataStoreMetaInfo
    TotalCountType uint8
}

type DataStoreSpecificMetaInfo struct {
    Base Structure
    DataId uint64
    OwnerId uint64
    Size uint32
    DataType uint16
    Version uint32
}

type DataStoreSpecificMetaInfoV1 struct {
    Base Structure
    DataId uint32
    OwnerId uint64
    Size uint32
    DataType uint16
    Version uint16
}

type DataStoreTouchObjectParam struct {
    Base Structure
    DataId uint64
    LockId uint32
    AccessPassword uint64
}

type DeletionEntry struct {
    Base Structure
    IdGathering uint32
    Pid PID
    UiReason uint32
}

type FindMatchmakeSessionByParticipantParam struct {
    PrincipalIdList []PID
    ResultOptions uint32
    BlockListParam MatchmakeBlockListParam
}

type FindMatchmakeSessionByParticipantResult struct {
    PrincipalId PID
    Session MatchmakeSession
}

type FriendData struct {
    Pid uint32
    StrName string
    ByRelationship uint8
    UiDetails uint32
    StrStatus string
}

type FriendInfo struct {
    NNAInfo NNAInfo
    NintendoPresence NintendoPresenceV2
    StatusMessage Comment
    BecameFriend DateTime
    LastOnline DateTime
    Unknown uint64
}

type FriendMii struct {
    Unknown uint32
    Mii Mii
    Unknown2 DateTime
}

type FriendMiiList struct {
    Unknown uint32
    MiiList MiiList
    Unknown2 DateTime
}

type FriendMiiRequest struct {
    Unknown uint32
    Unknown2 DateTime
}

type FriendPersistentInfo struct {
    Unknown uint32
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
    GameKey GameKey
    Unknown2 string
    Unknown3 DateTime
    Unknown4 DateTime
    Unknown5 DateTime
}

type FriendPicture struct {
    Unknown uint32
    Data Buffer
    DateTime DateTime
}

type FriendPresence struct {
    Unknown uint32
    NintendoPresence NintendoPresence
}

type FriendRelationship struct {
    Unknown uint32
    Unknown2 uint64
    Unknown3 uint8
}

type FriendRequest struct {
    PrincipalBasicInfo PrincipalBasicInfo
    Message FriendRequestMessage
    SentOn DateTime
}

type FriendRequestMessage struct {
    Unknown uint64
    Unknown2 uint8
    Unknown3 uint8
    Message string
    Unknown4 uint8
    Unknown5 string
    GameKey GameKey
    Unknown6 DateTime
    ExpiresOn DateTime
}

type GameKey struct {
    TitleId uint64
    TitleVersion uint16
}

type Gathering struct {
    Base Structure
    IdMyself uint32
    PidOwner PID
    PidHost PID
    UiMinParticipants uint16
    UiMaxParticipants uint16
    UiParticipationPolicy uint32
    UiPolicyArgument uint32
    UiFlags uint32
    UiState uint32
    StrDescription string
}

type GatheringStats struct {
    Base Structure
    PidParticipant uint32
    UiFlags uint32
    LstValues []float32
}

type GatheringURLs struct {
    Base Structure
    Gid uint32
    LstStationURLs []StationURL
}

type Invitation struct {
    Base Structure
    IdGathering uint32
    IdGuest uint32
    StrMessage string
}

type JoinMatchmakeSessionParam struct {
    Base Structure
    Gid uint32
    AdditionalParticipants []PID
    GidForParticipationCheck uint32
    JoinMatchmakeSessionOption uint32
    JoinMatchmakeSessionBehavior uint8
    StrUserPassword string
    StrSystemPassword string
    JoinMessage string
    ParticipationCount uint16
    ExtraParticipants uint16
    BlockListParam MatchmakeBlockListParam
}

type MatchmakeBlockListParam struct {
    Base Structure
    OptionFlag uint32
}

type MatchmakeParam struct {
    Base Structure
    Params map[string]Variant
}

type MatchmakeSession struct {
    Base Structure
    Base2 Gathering
    GameMode uint32
    Attribs []uint32
    OpenParticipation bool
    MatchmakeSystemType uint32
    ApplicationBuffer Buffer
    ParticipationCount uint32
    ProgressScore uint8
    SessionKey Buffer
    Option0 uint32
    MatchmakeParam MatchmakeParam
    StartedTime DateTime
    UserPassword string
    ReferGid uint32
    UserPasswordEnabled bool
    SystemPasswordEnabled bool
    Codeword string
}

type MatchmakeSessionSearchCriteria struct {
    Base Structure
    Attribs []string
    GameMode string
    MinParticipants string
    MaxParticipants string
    MatchmakeSystemType string
    VacantOnly bool
    ExcludeLocked bool
    ExcludeNonHostPid bool
    SelectionMethod uint32
    VacantParticipants uint16
    MatchmakeParam MatchmakeParam
    ExcludeUserPasswordSet bool
    ExcludeSystemPasswordSet bool
    ReferGid uint32
    Codeword string
    ResultRange ResultRange
}

type MessageRecipient struct {
    UiRecipientType uint32
    PrincipalId PID
    GatheringId uint32
}

type Mii struct {
    Unknown string
    Unknown2 bool
    Unknown3 uint8
    MiiData Buffer
}

type MiiList struct {
    Unknown string
    Unknown2 bool
    Unknown3 uint8
    MiiDataList []Buffer
}

type MiiV2 struct {
    Name string
    Unknown uint8
    Unknown2 uint8
    MiiDataFFLStoreData Buffer
    Unknown3 DateTime
}

type MyProfile struct {
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
    Unknown uint64
    Unknown2 string
    Unknown3 string
}

type NNAInfo struct {
    PrincipalBasicInfo PrincipalBasicInfo
    Unknown uint8
    Unknown2 uint8
}

type NintendoNotificationEvent struct {
    Base Structure
    UiType uint32
    PidSender PID
    Dataholder Data
}

type NintendoNotificationEventGeneral struct {
    Base Structure
    U32Param uint32
    U64Param1 uint64
    U64Param2 uint64
    StrParam string
}

type NintendoNotificationEventProfile struct {
    Base Structure
    Region uint8
    Country uint8
    Area uint8
    Language uint8
    Platform uint8
}

type NintendoPresence struct {
    ChangedBitFlag uint32
    GameKey GameKey
    GameModeDescription string
    JoinAvailabilityFlag uint32
    MatchmakeSystemType uint8
    JoinGameID uint32
    JoinGameMode uint32
    OwnerPrincipalID PID
    JoinGroupID uint32
    ApplicationArg Buffer
}

type NintendoPresenceV2 struct {
    ChangedFlags uint32
    IsOnline bool
    GameKey GameKey
    Unknown1 uint8
    Message string
    Unknown2 uint32
    Unknown3 uint8
    GameServerId uint32
    Unknown4 uint32
    Pid PID
    GatheringId uint32
    ApplicationData Buffer
    Unknown5 uint8
    Unknown6 uint8
    Unknown7 uint8
}

type NotificationEvent struct {
    Base Structure
    PidSource PID
    UiType uint32
    UiParam1 uint32
    UiParam2 uint32
    StrParam string
}

type ParticipantDetails struct {
    Base Structure
    IdParticipant uint32
    StrName string
    StrMessage string
    UiParticipants uint16
}

type PersistentGathering struct {
    Base Structure
    Base2 Gathering
    CommunityType uint32
    Password string
    Attribs []uint32
    ApplicationBuffer Buffer
    ParticipationStartDate DateTime
    ParticipationEndDate DateTime
    MatchmakeSessionCount uint32
    ParticipationCount uint32
}

type PersistentNotification struct {
    Unknown uint64
    Unknown2 uint32
    Unknown3 uint32
    Unknown4 uint32
    Unknown5 string
}

type PlayedGame struct {
    GameKey GameKey
    DateTime DateTime
}

type PlayingSession struct {
    Base Structure
    PrincipalId PID
    Gathering Data
}

type PrincipalBasicInfo struct {
    Pid PID
    NNID string
    Mii MiiV2
    Unknown uint8
}

type PrincipalPreference struct {
    Unknown bool
    Unknown2 bool
    Unknown3 bool
}

type PrincipalRequestBlockSetting struct {
    Unknown uint32
    Unknown2 bool
}

type RVConnectionData struct {
    UrlRegularProtocols string
    LstSpecialProtocols []uint8
    UrlSpecialProtocols string
}

type RankingCachedResult struct {
    CreatedTime DateTime
    ExpiredTime DateTime
    MaxLength uint8
}

type RankingChangeAttributesParam struct {
    Base Structure
    ModificationFlag uint8
    Groups []uint8
    Param uint64
}

type RankingOrderParam struct {
    Base Structure
    OrderCalculation uint8
    GroupIndex uint8
    GroupNum uint8
    TimeScope uint8
    Offset uint32
    Length uint8
}

type RankingRankData struct {
    Base Structure
    PrincipalId PID
    UniqueId uint64
    Order uint32
    Category uint32
    Score uint32
    Groups []uint8
    Param uint64
    CommonData Buffer
}

type RankingResult struct {
    Base Structure
    RankDataList []RankingRankData
    TotalCount uint32
    SinceTime DateTime
}

type RankingScoreData struct {
    Base Structure
    PrincipalId PID
    UniqueId uint64
    Order uint32
    Category uint32
    Score uint32
    Groups []uint8
    Param uint64
    CommonData Buffer
}

type RankingStats struct {
    Base Structure
    StatsList []float64
}

type RelationshipData struct {
    Pid uint32
    StrName string
    ByRelationship uint8
    UiDetails uint32
    ByStatus uint8
}

type ResultRange struct {
    Base Structure
    UiOffset uint32
    UiSize uint32
}

type SimpleCommunity struct {
    Base Structure
    GatheringID uint32
    MatchmakeSessionCount uint32
}

type SimplePlayingSession struct {
    Base Structure
    PrincipalID PID
    GatheringID uint32
    GameMode uint32
    Attribute_0 uint32
}

type Structure struct {
}

type UniqueIdInfo struct {
    Base Structure
    NexUniqueId uint64
    NexUniqueIdPassword uint64
}

type UpdateMatchmakeSessionParam struct {
    Base Structure
    Gid uint32
    ModificationFlag uint32
    Attributes []uint32
    OpenParticipation bool
    ApplicationBuffer Buffer
    ProgressScore uint8
    MatchmakeParam MatchmakeParam
    StartedTime DateTime
    UserPassword string
    GameMode uint32
    Description string
    MinParticipants uint16
    MaxParticipants uint16
    MatchmakeSystemType uint32
    ParticipationPolicy uint32
    PolicyArgument uint32
    Codeword string
}

type UserMessage struct {
    UiID uint32
    UiParentID uint32
    PidSender PID
    Receptiontime DateTime
    UiLifeTime uint32
    UiFlags uint32
    StrSubject string
    StrSender string
    MessageRecipient MessageRecipient
}

