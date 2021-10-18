
CREATE TABLE `companies` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `rut` int(10) DEFAULT NULL,
  `dv` char(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `razon` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `city` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `commune` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `people_id` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`rut`, `dv`, `razon`, `address`, `city`, `commune`, `people_id`) VALUES
( 22221960, '4', 'COMPANIA', 'TENO 123', 'CURICO', 'TENO', 1);

-- --------------------------------------------------------

--
-- Table structure for table `people`
--

CREATE TABLE `people` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `rut` int(10) DEFAULT NULL,
  `dv` char(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lastname` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `city` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `commune` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(15) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profile_id` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `profile`
--

CREATE TABLE `profile` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `profile` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `profile`
--

INSERT INTO `profile` ( `profile`) VALUES
( 'ADMINISTRADOR'),
('CLIENTE'),
('REPARTIDOR');

-- --------------------------------------------------------

--
-- Table structure for table `rates`
--

CREATE TABLE `rates` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `from` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `to` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` double DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `rol` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `rol`) VALUES
( 'ADMINISTRADOR'),
( 'CLIENTE'),
( 'REPARTIDOR');

-- --------------------------------------------------------

--
-- Table structure for table `shipping`
--

CREATE TABLE `shipping` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `order_nro` int(11) DEFAULT NULL,
  `boleta_nro` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` double DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  `shipping_states_id` int(11) DEFAULT NULL,
  `companies_id` int(11) DEFAULT NULL,
  `delivery_id` int(11) DEFAULT NULL,
  `client_id` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `shipping_states`
--

CREATE TABLE `shipping_states` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `state` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `user` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NOT NULL UNIQUE,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rol_id` int(11) DEFAULT NULL,
  `user_state_id` int(11) DEFAULT NULL,
  `people_id` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user_state`
--

CREATE TABLE `user_state` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT,
  `state` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `user_state`
--

INSERT INTO `user_state` (`state`) VALUES
('ACTIVO'),
( 'SUSPENDIDO'),
( 'BLOQUEADO'),
('ELIMINADO');


COMMIT;