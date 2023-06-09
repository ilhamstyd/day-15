PGDMP                         {         
   db_project    10.23    15.2     �
           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �
           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �
           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                        1262    24674 
   db_project    DATABASE     �   CREATE DATABASE db_project WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE db_project;
                postgres    false                        2615    2200    public    SCHEMA     2   -- *not* creating schema, since initdb creates it
 2   -- *not* dropping schema, since initdb creates it
                postgres    false                       0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                   postgres    false    6            �            1259    24677    tb_projects    TABLE     !  CREATE TABLE public.tb_projects (
    id integer NOT NULL,
    description character varying(255) NOT NULL,
    image character varying(255) NOT NULL,
    author integer,
    name character varying(255) NOT NULL,
    start_date character varying(10),
    end_date character varying(10)
);
    DROP TABLE public.tb_projects;
       public            postgres    false    6            �            1259    24675    tb_projects_id_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.tb_projects_id_seq;
       public          postgres    false    6    197                       0    0    tb_projects_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.tb_projects_id_seq OWNED BY public.tb_projects.id;
          public          postgres    false    196            �            1259    24688    tb_user    TABLE     �   CREATE TABLE public.tb_user (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);
    DROP TABLE public.tb_user;
       public            postgres    false    6            �            1259    24686    tb_user_id_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.tb_user_id_seq;
       public          postgres    false    199    6                       0    0    tb_user_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.tb_user_id_seq OWNED BY public.tb_user.id;
          public          postgres    false    198            v
           2604    24680    tb_projects id    DEFAULT     p   ALTER TABLE ONLY public.tb_projects ALTER COLUMN id SET DEFAULT nextval('public.tb_projects_id_seq'::regclass);
 =   ALTER TABLE public.tb_projects ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    197    196    197            w
           2604    24691 
   tb_user id    DEFAULT     h   ALTER TABLE ONLY public.tb_user ALTER COLUMN id SET DEFAULT nextval('public.tb_user_id_seq'::regclass);
 9   ALTER TABLE public.tb_user ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    199    198    199            �
          0    24677    tb_projects 
   TABLE DATA           a   COPY public.tb_projects (id, description, image, author, name, start_date, end_date) FROM stdin;
    public          postgres    false    197   `       �
          0    24688    tb_user 
   TABLE DATA           <   COPY public.tb_user (id, name, email, password) FROM stdin;
    public          postgres    false    199   �                  0    0    tb_projects_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.tb_projects_id_seq', 7, true);
          public          postgres    false    196                       0    0    tb_user_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.tb_user_id_seq', 1, true);
          public          postgres    false    198            z
           2606    24685    tb_projects tb_projects_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.tb_projects
    ADD CONSTRAINT tb_projects_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.tb_projects DROP CONSTRAINT tb_projects_pkey;
       public            postgres    false    197            |
           2606    24696    tb_user tb_user_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.tb_user
    ADD CONSTRAINT tb_user_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.tb_user DROP CONSTRAINT tb_user_pkey;
       public            postgres    false    199            x
           1259    24702    fki_fk_projects_author    INDEX     P   CREATE INDEX fki_fk_projects_author ON public.tb_projects USING btree (author);
 *   DROP INDEX public.fki_fk_projects_author;
       public            postgres    false    197            }
           2606    24697    tb_projects fk_projects_author    FK CONSTRAINT     �   ALTER TABLE ONLY public.tb_projects
    ADD CONSTRAINT fk_projects_author FOREIGN KEY (author) REFERENCES public.tb_user(id) NOT VALID;
 H   ALTER TABLE ONLY public.tb_projects DROP CONSTRAINT fk_projects_author;
       public          postgres    false    2684    197    199            �
   |   x�]�M� �u{
.�A�0n:�#j,��?�8�״�Kk��e::b�� `0#�%@�Yp!�
�E������6�6ϖK�h2�N�7)�O������������>G�G�.i�}��'�;�      �
   (   x�3����H̅�ũ%��)Y�z�����F�&\1z\\\ ��
�     